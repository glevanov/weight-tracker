package tests

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"testing"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type AddWeightRequest struct {
	Weight string `json:"weight"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type WeightResponse struct {
	Weight    float64   `json:"weight"`
	Timestamp time.Time `json:"timestamp"`
}

type SuccessResponse[T any] struct {
	IsSuccess bool `json:"isSuccess"`
	Data      T    `json:"data,omitempty"`
}

type ErrorResponse struct {
	IsSuccess bool   `json:"isSuccess"`
	Error     string `json:"error,omitempty"`
}

// createTestUser creates a test user directly in Postgres
// Password: "testpassword", Salt: "0102030405060708090a0b0c0d0e0f10"
func createTestUser(databaseURL string) error {
	ctx := context.Background()
	pool, err := pgxpool.New(ctx, databaseURL)
	if err != nil {
		return err
	}
	defer pool.Close()

	_, err = pool.Exec(ctx, `
		CREATE TABLE IF NOT EXISTS users (
			username TEXT PRIMARY KEY,
			password TEXT NOT NULL,
			salt TEXT NOT NULL
		)
	`)
	if err != nil {
		return err
	}

	_, err = pool.Exec(ctx, `
		CREATE TABLE IF NOT EXISTS weights (
			id BIGSERIAL PRIMARY KEY,
			username TEXT NOT NULL,
			weight DOUBLE PRECISION NOT NULL,
			timestamp TIMESTAMPTZ NOT NULL DEFAULT NOW()
		)
	`)
	if err != nil {
		return err
	}

	// Hash generated using scrypt with N=16384, r=8, p=1, keyLen=64
	_, err = pool.Exec(ctx,
		"INSERT INTO users (username, password, salt) VALUES ($1, $2, $3)",
		"testuser",
		"404ba06bdb03dc9a8a9ad7ea8e1f13a58d0c4a2a600580bf9ac558147c20afd960e7300e8ce8d0874dbd6be8cf4147caf07182787e468001f06d17df9b7e42b5",
		"0102030405060708090a0b0c0d0e0f10",
	)

	return err
}

func login(baseURL, username, password string) (string, error) {
	reqBody := LoginRequest{Username: username, Password: password}
	jsonBody, err := json.Marshal(reqBody)
	if err != nil {
		return "", err
	}

	resp, err := http.Post(
		baseURL+"/login",
		"application/json",
		bytes.NewBuffer(jsonBody),
	)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var loginResp SuccessResponse[string]
	if err := json.Unmarshal(body, &loginResp); err != nil {
		return "", err
	}

	return loginResp.Data, nil
}

func TestIntegration(t *testing.T) {
	baseURL, databaseURL, cleanup := SetupTestEnvironment(t)
	defer cleanup()

	err := createTestUser(databaseURL)
	require.NoError(t, err)

	var token string
	var weights []WeightResponse

	t.Run("HealthCheck", func(t *testing.T) {
		var resp *http.Response
		var err error
		for range 10 {
			resp, err = http.Get(baseURL + "/health-check")
			if err == nil {
				break
			}
			time.Sleep(time.Second)
		}
		require.NoError(t, err)
		require.NotNil(t, resp)
		defer resp.Body.Close()

		assert.Equal(t, http.StatusOK, resp.StatusCode)
		assert.Equal(t, "application/json", resp.Header.Get("Content-Type"))

		body, err := io.ReadAll(resp.Body)
		require.NoError(t, err)

		var healthResp SuccessResponse[string]
		err = json.Unmarshal(body, &healthResp)
		require.NoError(t, err)

		assert.True(t, healthResp.IsSuccess)
		assert.Equal(t, "OK", healthResp.Data)
	})

	t.Run("Login", func(t *testing.T) {
		var err error
		token, err = login(baseURL, "testuser", "testpassword")
		require.NoError(t, err)
		assert.NotEmpty(t, token)
	})

	t.Run("SessionCheck", func(t *testing.T) {
		req, err := http.NewRequest("GET", baseURL+"/session-check", nil)
		require.NoError(t, err)
		req.Header.Set("Authorization", "Bearer "+token)

		resp, err := http.DefaultClient.Do(req)
		require.NoError(t, err)
		defer resp.Body.Close()

		assert.Equal(t, http.StatusOK, resp.StatusCode)

		body, err := io.ReadAll(resp.Body)
		require.NoError(t, err)

		var sessionResp SuccessResponse[string]
		err = json.Unmarshal(body, &sessionResp)
		require.NoError(t, err)

		assert.True(t, sessionResp.IsSuccess)
		assert.Equal(t, "OK", sessionResp.Data)
	})

	t.Run("AddWeightWithoutAuth", func(t *testing.T) {
		reqBody := AddWeightRequest{Weight: "82.5"}
		jsonBody, err := json.Marshal(reqBody)
		require.NoError(t, err)

		resp, err := http.Post(
			baseURL+"/weights",
			"application/json",
			bytes.NewBuffer(jsonBody),
		)
		require.NoError(t, err)
		defer resp.Body.Close()

		assert.Equal(t, http.StatusUnauthorized, resp.StatusCode)
	})

	t.Run("AddWeight", func(t *testing.T) {
		reqBody := AddWeightRequest{Weight: "82.5"}
		jsonBody, err := json.Marshal(reqBody)
		require.NoError(t, err)

		req, err := http.NewRequest("POST", baseURL+"/weights", bytes.NewBuffer(jsonBody))
		require.NoError(t, err)
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", "Bearer "+token)

		resp, err := http.DefaultClient.Do(req)
		require.NoError(t, err)
		defer resp.Body.Close()

		assert.Equal(t, http.StatusCreated, resp.StatusCode)
		assert.Equal(t, "application/json", resp.Header.Get("Content-Type"))

		body, err := io.ReadAll(resp.Body)
		require.NoError(t, err)

		var successResp SuccessResponse[string]
		err = json.Unmarshal(body, &successResp)
		require.NoError(t, err)

		assert.True(t, successResp.IsSuccess)
		assert.NotEmpty(t, successResp.Data)
	})

	t.Run("AddAnotherWeight", func(t *testing.T) {
		// Small delay to ensure different timestamps
		time.Sleep(100 * time.Millisecond)

		reqBody := AddWeightRequest{Weight: "83.1"}
		jsonBody, err := json.Marshal(reqBody)
		require.NoError(t, err)

		req, err := http.NewRequest("POST", baseURL+"/weights", bytes.NewBuffer(jsonBody))
		require.NoError(t, err)
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", "Bearer "+token)

		resp, err := http.DefaultClient.Do(req)
		require.NoError(t, err)
		defer resp.Body.Close()

		assert.Equal(t, http.StatusCreated, resp.StatusCode)
		assert.Equal(t, "application/json", resp.Header.Get("Content-Type"))

		body, err := io.ReadAll(resp.Body)
		require.NoError(t, err)

		var successResp SuccessResponse[string]
		err = json.Unmarshal(body, &successResp)
		require.NoError(t, err)

		assert.True(t, successResp.IsSuccess)
		assert.NotEmpty(t, successResp.Data)
	})

	t.Run("GetWeights", func(t *testing.T) {
		req, err := http.NewRequest("GET", baseURL+"/weights", nil)
		require.NoError(t, err)
		req.Header.Set("Authorization", "Bearer "+token)

		resp, err := http.DefaultClient.Do(req)
		require.NoError(t, err)
		defer resp.Body.Close()

		assert.Equal(t, http.StatusOK, resp.StatusCode)
		assert.Equal(t, "application/json", resp.Header.Get("Content-Type"))

		body, err := io.ReadAll(resp.Body)
		require.NoError(t, err)

		var successResp SuccessResponse[[]WeightResponse]
		err = json.Unmarshal(body, &successResp)
		require.NoError(t, err)

		assert.True(t, successResp.IsSuccess)

		weights = successResp.Data

		require.Len(t, weights, 2)

		assert.Equal(t, 82.5, weights[0].Weight)
		assert.False(t, weights[0].Timestamp.IsZero())

		assert.Equal(t, 83.1, weights[1].Weight)
		assert.False(t, weights[1].Timestamp.IsZero())

		assert.True(t, weights[0].Timestamp.Before(weights[1].Timestamp) ||
			weights[0].Timestamp.Equal(weights[1].Timestamp),
			"Weights should be sorted by timestamp ascending")
	})
}

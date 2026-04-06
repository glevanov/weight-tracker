package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"time"

	"weight-tracker-service/internal/auth"
	"weight-tracker-service/internal/database"
	"weight-tracker-service/internal/i18n"
	"weight-tracker-service/internal/logger"
	"weight-tracker-service/internal/validation"
)

type Weight struct {
	Weight    float64   `json:"weight"`
	Timestamp time.Time `json:"timestamp"`
}

type AddWeightRequest struct {
	Weight string `json:"weight"`
}

func GetWeights(w http.ResponseWriter, r *http.Request) {
	lang := i18n.ExtractLang(r)

	username := auth.UsernameFromContext(r.Context())
	logger.Info("get weights request", "username", username)

	startStr := r.URL.Query().Get("start")
	endStr := r.URL.Query().Get("end")

	queryParts := []string{"SELECT weight, timestamp FROM weights WHERE username = $1"}
	args := []any{username}

	if startStr != "" {
		startTime, err := time.Parse(time.RFC3339, startStr)
		if err != nil {
			logger.Warn("get weights failed: invalid start timestamp", "username", username, "start", startStr)
			writeError(w, http.StatusBadRequest, i18n.Translate(lang, validation.ErrTimestampNotDate))
			return
		}
		args = append(args, startTime)
		queryParts = append(queryParts, "AND timestamp >= $"+strconv.Itoa(len(args)))
	}

	if endStr != "" {
		endTime, err := time.Parse(time.RFC3339, endStr)
		if err != nil {
			logger.Warn("get weights failed: invalid end timestamp", "username", username, "end", endStr)
			writeError(w, http.StatusBadRequest, i18n.Translate(lang, validation.ErrTimestampNotDate))
			return
		}
		args = append(args, endTime)
		queryParts = append(queryParts, "AND timestamp <= $"+strconv.Itoa(len(args)))
	}

	queryParts = append(queryParts, "ORDER BY timestamp ASC")
	query := strings.Join(queryParts, " ")

	pool := database.GetPool()
	rows, err := pool.Query(r.Context(), query, args...)
	if err != nil {
		logger.Error("get weights failed: database query error", "username", username, "error", err)
		writeError(w, http.StatusInternalServerError, i18n.Translate(lang, validation.ErrUnknown))
		return
	}
	defer rows.Close()

	weights := make([]Weight, 0)
	for rows.Next() {
		var weight Weight
		if err := rows.Scan(&weight.Weight, &weight.Timestamp); err != nil {
			logger.Error("get weights failed: row scan error", "username", username, "error", err)
			writeError(w, http.StatusInternalServerError, i18n.Translate(lang, validation.ErrUnknown))
			return
		}
		weights = append(weights, weight)
	}

	if err := rows.Err(); err != nil {
		logger.Error("get weights failed: cursor decode error", "username", username, "error", err)
		writeError(w, http.StatusInternalServerError, i18n.Translate(lang, validation.ErrUnknown))
		return
	}

	logger.Info("get weights success", "username", username, "count", len(weights))
	writeSuccess(w, http.StatusOK, weights)
}

func AddWeight(w http.ResponseWriter, r *http.Request) {
	lang := i18n.ExtractLang(r)

	username := auth.UsernameFromContext(r.Context())
	logger.Info("add weight request", "username", username)

	var req AddWeightRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		logger.Warn("add weight failed: failed to parse request", "username", username, "error", err)
		writeError(w, http.StatusBadRequest, i18n.Translate(lang, validation.ErrWeightFailedToParse))
		return
	}

	weight, errMsg := validation.ValidateAndFormatWeight(req.Weight)
	if errMsg != "" {
		logger.Warn("add weight failed: validation error", "username", username, "weight", req.Weight, "error", errMsg)
		writeError(w, http.StatusBadRequest, i18n.Translate(lang, errMsg))
		return
	}

	pool := database.GetPool()
	_, err := pool.Exec(
		r.Context(),
		"INSERT INTO weights (username, weight, timestamp) VALUES ($1, $2, $3)",
		username,
		weight,
		time.Now(),
	)
	if err != nil {
		logger.Error("add weight failed: database insert error", "username", username, "error", err)
		writeError(w, http.StatusInternalServerError, i18n.Translate(lang, validation.ErrUnknown))
		return
	}

	logger.Info("add weight success", "username", username, "weight", weight)
	writeSuccess(w, http.StatusCreated, i18n.Translate(lang, validation.ResponseWeightAdded))
}

func writeSuccess(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(SuccessResult{
		IsSuccess: true,
		Data:      data,
	})
}

func writeError(w http.ResponseWriter, status int, err string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(ErrorResult{
		IsSuccess: false,
		Error:     err,
	})
}

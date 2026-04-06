package tests

import (
	"context"
	"fmt"
	"path/filepath"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/network"
	"github.com/testcontainers/testcontainers-go/wait"
)

const (
	TestDBName = "weight-tracker-test"
)

func SetupTestEnvironment(t *testing.T) (string, string, func()) {
	ctx := context.Background()

	// Create a Docker network for container communication
	nw, err := network.New(ctx)
	require.NoError(t, err, "Failed to create test network")
	require.NotNil(t, nw, "Network should not be nil")

	postgresContainer, err := setupPostgresContainer(t, nw.Name)
	require.NoError(t, err, "Failed to start Postgres container")

	postgresHost, err := postgresContainer.Host(ctx)
	require.NoError(t, err, "Failed to get Postgres host")

	postgresPort, err := postgresContainer.MappedPort(ctx, "5432")
	require.NoError(t, err, "Failed to get Postgres port")

	databaseURL := fmt.Sprintf("postgres://postgres:postgres@%s:%s/%s?sslmode=disable", postgresHost, postgresPort.Port(), TestDBName)

	baseURL, serviceCleanup := setupServiceContainer(t, nw.Name)

	cleanup := func() {
		serviceCleanup()
		if err := postgresContainer.Terminate(ctx); err != nil {
			t.Logf("Failed to terminate Postgres container: %v", err)
		}
		if err := nw.Remove(ctx); err != nil {
			t.Logf("Failed to remove test network: %v", err)
		}
	}

	return baseURL, databaseURL, cleanup
}

func setupPostgresContainer(t *testing.T, networkName string) (testcontainers.Container, error) {
	ctx := context.Background()

	req := testcontainers.ContainerRequest{
		Image:        "postgres:17-alpine",
		ExposedPorts: []string{"5432/tcp"},
		Networks:     []string{networkName},
		NetworkAliases: map[string][]string{
			networkName: {"postgres"},
		},
		Env: map[string]string{
			"POSTGRES_USER":     "postgres",
			"POSTGRES_PASSWORD": "postgres",
			"POSTGRES_DB":       TestDBName,
		},
		WaitingFor: wait.ForListeningPort("5432/tcp").WithStartupTimeout(60 * time.Second),
	}

	container, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
		Reuse:            false,
	})
	if err != nil {
		return nil, err
	}

	return container, nil
}

func setupServiceContainer(t *testing.T, networkName string) (string, func()) {
	ctx := context.Background()

	projectRoot, err := filepath.Abs("..")
	require.NoError(t, err, "Failed to get project root")

	serviceDatabaseURL := fmt.Sprintf("postgres://postgres:postgres@postgres:5432/%s?sslmode=disable", TestDBName)

	req := testcontainers.ContainerRequest{
		FromDockerfile: testcontainers.FromDockerfile{
			Context:    projectRoot,
			Dockerfile: "Dockerfile.test",
		},
		ExposedPorts: []string{"3000/tcp"},
		Networks:     []string{networkName},
		WaitingFor:   wait.ForListeningPort("3000/tcp").WithStartupTimeout(60 * time.Second),
		Env: map[string]string{
			"PORT":         "3000",
			"FRONTEND_URL": "http://localhost:5173",
			"DATABASE_URL": serviceDatabaseURL,
			"JWT_SECRET":   "test-jwt-secret-key",
		},
	}

	container, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
		Reuse:            false,
	})
	require.NoError(t, err, "Failed to start service container")
	require.NotNil(t, container, "Container should not be nil")

	mappedPort, err := container.MappedPort(ctx, "3000")
	require.NoError(t, err, "Failed to get mapped port")

	host, err := container.Host(ctx)
	require.NoError(t, err, "Failed to get container host")

	baseURL := fmt.Sprintf("http://%s:%s", host, mappedPort.Port())

	cleanup := func() {
		if err := container.Terminate(ctx); err != nil {
			t.Logf("Failed to terminate service container: %v", err)
		}
	}

	return baseURL, cleanup
}

package config

import (
	"fmt"
	"net/url"
	"os"
	"strconv"
	"time"
)

type Config struct {
	Port            int
	FrontendURL     string
	DatabaseURL     string
	JWTSecret       string
	SessionDuration time.Duration
}

func Load() *Config {
	port := 3000
	if p := envFirst("WEITHG_TRACKER_PORT", "PORT"); p != "" {
		if parsed, err := strconv.Atoi(p); err == nil {
			port = parsed
		}
	}

	frontendURL := envFirst("WEITHG_TRACKER_FRONTEND_URL", "FRONTEND_URL")

	databaseURL := envFirst(
		"WEITHG_TRACKER_DATABASE_URL",
		"DATABASE_URL",
		"CONNECTION_URI",
	)
	if databaseURL == "" {
		databaseURL = databaseURLFromParts()
	}
	if databaseURL == "" {
		panic(fmt.Errorf("WEITHG_TRACKER_DATABASE_URL or database connection parts are required"))
	}

	jwtSecret := envFirst("WEITHG_TRACKER_JWT_SECRET", "JWT_SECRET")
	if jwtSecret == "" {
		panic(fmt.Errorf("WEITHG_TRACKER_JWT_SECRET environment variable is required"))
	}

	// 5 years session duration
	sessionDuration := 5 * 365 * 24 * time.Hour

	return &Config{
		Port:            port,
		FrontendURL:     frontendURL,
		DatabaseURL:     databaseURL,
		JWTSecret:       jwtSecret,
		SessionDuration: sessionDuration,
	}
}

func envFirst(keys ...string) string {
	for _, key := range keys {
		if value := os.Getenv(key); value != "" {
			return value
		}
	}

	return ""
}

func databaseURLFromParts() string {
	host := envFirst("WEITHG_TRACKER_DB_HOST")
	if host == "" {
		host = "127.0.0.1"
	}

	port := envFirst("WEITHG_TRACKER_DB_PORT")
	if port == "" {
		port = "5432"
	}

	sslMode := envFirst("WEITHG_TRACKER_DB_SSLMODE")
	if sslMode == "" {
		sslMode = "disable"
	}

	databaseName := envFirst("WEITHG_TRACKER_DB_NAME")
	user := envFirst("WEITHG_TRACKER_DB_USER")
	password := envFirst("WEITHG_TRACKER_DB_PASSWORD")

	if databaseName == "" || user == "" || password == "" {
		return ""
	}

	return fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s",
		url.QueryEscape(user),
		url.QueryEscape(password),
		host,
		port,
		databaseName,
		sslMode,
	)
}

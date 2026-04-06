package config

import (
	"fmt"
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
	if p := os.Getenv("PORT"); p != "" {
		if parsed, err := strconv.Atoi(p); err == nil {
			port = parsed
		}
	}

	frontendURL := os.Getenv("FRONTEND_URL")
	if frontendURL == "" {
		frontendURL = "http://localhost:5173"
	}

	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		databaseURL = os.Getenv("CONNECTION_URI")
	}
	if databaseURL == "" {
		panic(fmt.Errorf("DATABASE_URL environment variable is required"))
	}

	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		panic(fmt.Errorf("JWT_SECRET environment variable is required"))
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

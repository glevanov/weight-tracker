package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"weight-tracker-service/internal/auth"
	"weight-tracker-service/internal/config"
	"weight-tracker-service/internal/database"
	"weight-tracker-service/internal/handlers"
	"weight-tracker-service/internal/logger"
	"weight-tracker-service/internal/static"
)

func main() {
	cfg := config.Load()

	if err := database.Connect(cfg.DatabaseURL); err != nil {
		logger.Error("database connection error", "error", err)
		return
	}
	defer database.Disconnect()

	r := chi.NewRouter()
	if cfg.FrontendURL != "" {
		r.Use(cors.Handler(cors.Options{
			AllowedOrigins:   []string{cfg.FrontendURL},
			AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
			AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type"},
			AllowCredentials: true,
		}))
	}

	r.Route("/api", func(r chi.Router) {
		r.Get("/health-check", handlers.HealthCheck)
		r.Post("/login", handlers.Login(cfg))

		r.Group(func(r chi.Router) {
			r.Use(auth.Middleware(cfg.JWTSecret))
			r.Get("/weights", handlers.GetWeights)
			r.Post("/weights", handlers.AddWeight)
			r.Get("/session-check", handlers.SessionCheck)
		})
	})

	if staticHandler, ok := static.Handler(); ok {
		r.NotFound(func(w http.ResponseWriter, req *http.Request) {
			if strings.HasPrefix(req.URL.Path, "/api/") {
				http.NotFound(w, req)
				return
			}

			staticHandler.ServeHTTP(w, req)
		})
	}

	addr := fmt.Sprintf(":%d", cfg.Port)
	logger.Info("server starting", "port", cfg.Port)

	if err := http.ListenAndServe(addr, r); err != nil {
		logger.Error("server error", "error", err)
	}
}

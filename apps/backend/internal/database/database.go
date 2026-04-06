package database

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"weight-tracker-service/internal/logger"
)

var (
	pool *pgxpool.Pool
)

func Connect(uri string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	p, err := pgxpool.New(ctx, uri)
	if err != nil {
		logger.Error("database connection failed", "error", err)
		return err
	}

	if err := p.Ping(ctx); err != nil {
		p.Close()
		logger.Error("database ping failed", "error", err)
		return err
	}

	pool = p

	logger.Info("database connected")
	return nil
}

func Disconnect() {
	if pool != nil {
		pool.Close()
		logger.Info("database disconnected")
	}
}

func GetPool() *pgxpool.Pool {
	return pool
}

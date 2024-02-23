package db

import (
	"context"
	"github.com/jackc/pgx/v5"
	"rinha/backend/api/config"
)

func Connect(cfg *config.Config) (*pgx.Conn, error) {
	conn, err := pgx.Connect(context.Background(), cfg.DatabaseUrl)
	if err != nil {
		return conn, err
	}
	return conn, nil
}

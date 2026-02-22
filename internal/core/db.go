package core

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

func newDB(cfg *Config) (*pgxpool.Pool, error) {
	return pgxpool.New(context.Background(), cfg.DatabaseUrl)
}

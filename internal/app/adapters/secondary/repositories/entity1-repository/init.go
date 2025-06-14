package entity1_repository

import (
	"log/slog"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Entity1Repository struct {
	logger *slog.Logger
	config Config
	db     *pgxpool.Pool
}

type Config struct{}

func New(l *slog.Logger, cfg Config, db *pgxpool.Pool) *Entity1Repository {
	return &Entity1Repository{
		logger: l,
		config: cfg,
		db:     db,
	}
}

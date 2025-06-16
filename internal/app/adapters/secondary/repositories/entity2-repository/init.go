package entity2_repository

import (
	"log/slog"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Entity2Repository struct {
	logger *slog.Logger
	config Config
	db     *pgxpool.Pool
}
type Config struct{}

func New(l *slog.Logger, cfg Config, db *pgxpool.Pool) *Entity2Repository {
	return &Entity2Repository{
		logger: l,
		config: cfg,
		db:     db,
	}
}

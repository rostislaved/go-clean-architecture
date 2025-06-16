package entity4_repository

import (
	"log/slog"

	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
)

type Entity4Repository struct {
	logger *slog.Logger
	config Config
	db     *driver.Conn
}

type Config struct{}

func New(l *slog.Logger, cfg Config, db *driver.Conn) *Entity4Repository {
	return &Entity4Repository{
		logger: l,
		config: cfg,
		db:     db,
	}
}

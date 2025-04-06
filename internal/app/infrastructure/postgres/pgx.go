package postgres

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

func Pgx(l *slog.Logger, cfg Config) (*pgxpool.Pool, error) {
	connectingString, connectedString := getLogString(cfg)

	l.Info(connectingString)

	poolConfig, err := buildPgxPoolConfig(cfg)
	if err != nil {
		return nil, fmt.Errorf("failed to create config: %w", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	db, err := pgxpool.NewWithConfig(ctx, poolConfig)
	if err != nil {
		return nil, err
	}

	l.Info(connectedString)

	return db, nil
}

func buildPgxPoolConfig(cfg Config) (*pgxpool.Config, error) {
	connectionString := postgresConnectionString(cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Name, cfg.TimeZone)

	poolConfig, err := pgxpool.ParseConfig(connectionString)
	if err != nil {
		return nil, fmt.Errorf("failed to parse config: %w", err)
	}

	if cfg.MaxOpenConns != nil {
		poolConfig.MaxConns = int32(*cfg.MaxOpenConns) //nolint:gosec // won’t overflow for typical values
	}

	if cfg.ConnMaxIdleTime != nil {
		poolConfig.MaxConnIdleTime = *cfg.ConnMaxIdleTime
	}

	if cfg.ConnMaxLifetime != nil {
		poolConfig.MaxConnLifetime = *cfg.ConnMaxLifetime
	}

	return poolConfig, nil
}

package postgres

import (
	"context"
	"log/slog"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib" // db driver
	"github.com/jmoiron/sqlx"
)

const pgxDriverName = "pgx"

func Sqlx(l *slog.Logger, cfg Config) (*sqlx.DB, error) {
	connectingString, connectedString := getLogString(cfg)

	l.Info(connectingString)

	connectionString := postgresConnectionString(cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Name, cfg.TimeZone)

	db, err := sqlx.Open(pgxDriverName, connectionString)
	if err != nil {
		return nil, err
	}

	applyDBConfig(db, cfg)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}

	l.Info(connectedString)

	return db, nil
}

func applyDBConfig(db *sqlx.DB, cfg Config) {
	if cfg.MaxOpenConns != nil {
		db.SetMaxOpenConns(*cfg.MaxOpenConns)
	}

	if cfg.MaxIdleConns != nil {
		db.SetMaxIdleConns(*cfg.MaxIdleConns)
	}

	if cfg.ConnMaxIdleTime != nil {
		db.SetConnMaxIdleTime(*cfg.ConnMaxIdleTime)
	}

	if cfg.ConnMaxLifetime != nil {
		db.SetConnMaxLifetime(*cfg.ConnMaxLifetime)
	}
}

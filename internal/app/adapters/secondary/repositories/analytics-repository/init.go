package torgRepository

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"log/slog"
	"os"
	"time"

	_ "github.com/mailru/go-clickhouse/v2"

	"github.com/rostislaved/go-clean-architecture/internal/app/domain/config"
	"github.com/rostislaved/go-clean-architecture/internal/pkg/helpers"
	"github.com/rostislaved/go-clean-architecture/internal/pkg/repohelpers"
)

type AnalyticsRepository struct {
	logger *slog.Logger
	config config.DatabaseRelational
	DB     *sql.DB
}

func New(l *slog.Logger, cfg config.DatabaseRelational) *AnalyticsRepository {
	currentHostString := fmt.Sprintf("DB host: [%s:%s].", cfg.Host, cfg.Port)

	log.Println(currentHostString + " Подключение...")
	l.Info(currentHostString+" Подключение...", helpers.GetFunctionName())

	connectionString := repohelpers.GetConnectionString(cfg.Type, cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Name)

	db, err := sql.Open(cfg.Type, connectionString)
	if err != nil {
		l.Error(err.Error(), helpers.GetFunctionName())

		os.Exit(1)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		log.Println(err.Error(), helpers.GetFunctionName())

		l.Error(err.Error(), helpers.GetFunctionName())

		os.Exit(1)
	}

	log.Println(currentHostString + " Подключено!")
	l.Info(currentHostString+" Подключено!", helpers.GetFunctionName())

	return &AnalyticsRepository{
		logger: l,
		config: cfg,
		DB:     db,
	}
}

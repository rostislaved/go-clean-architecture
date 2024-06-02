package booksRepository

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"os"
	"time"

	sql "github.com/jmoiron/sqlx"

	_ "github.com/denisenkom/go-mssqldb"
	_ "github.com/lib/pq"

	"github.com/rostislaved/go-clean-architecture/internal/pkg/helpers"
	"github.com/rostislaved/go-clean-architecture/internal/pkg/repohelpers"
)

type FirstRepository struct {
	logger *slog.Logger
	config DatabaseRelational
	DB     *sql.DB
}

func New(l *slog.Logger, cfg DatabaseRelational) *FirstRepository {
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

	return &FirstRepository{
		logger: l,
		config: cfg,
		DB:     db,
	}
}

type DatabaseRelational struct {
	Type       string
	Host       string `config:"envVar"`
	Port       string `config:"envVar"`
	User       string `config:"envVar"`
	Password   string `config:"envVar"`
	Name       string
	Procedures map[string]string
}

package books_repository_clickhouse

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"log/slog"
	"time"

	_ "github.com/mailru/go-clickhouse/v2"

	"github.com/rostislaved/go-clean-architecture/internal/libs/helpers"
	"github.com/rostislaved/go-clean-architecture/internal/libs/repo-helpers"
)

type BooksRepositoryClickhouse struct {
	logger *slog.Logger
	config Config
	DB     *sql.DB
}

type Config struct {
	Type       string
	Host       string `config:"envVar"`
	Port       string `config:"envVar"`
	User       string `config:"envVar"`
	Password   string `config:"envVar"`
	Name       string
	Procedures map[string]string
}

func New(l *slog.Logger, cfg Config) *BooksRepositoryClickhouse {
	currentHostString := fmt.Sprintf("DB host: [%s:%s].", cfg.Host, cfg.Port)

	log.Println(currentHostString + " Подключение...")
	l.Info(currentHostString+" Подключение...", "source", helpers.GetFunctionName())

	connectionString := repo_helpers.GetConnectionString(cfg.Type, cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Name)

	db, err := sql.Open(cfg.Type, connectionString)
	if err != nil {
		l.Error(err.Error(), "source", helpers.GetFunctionName())

		panic(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		l.Error(err.Error(), "source", helpers.GetFunctionName())

		panic(err)
	}

	log.Println(currentHostString + " Подключено!")
	l.Info(currentHostString+" Подключено!", "source", helpers.GetFunctionName())

	return &BooksRepositoryClickhouse{
		logger: l,
		config: cfg,
		DB:     db,
	}
}

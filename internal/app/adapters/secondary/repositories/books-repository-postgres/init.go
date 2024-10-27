package books_repository_postgres

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"time"

	_ "github.com/denisenkom/go-mssqldb"
	sql "github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	"github.com/rostislaved/go-clean-architecture/internal/libs/helpers"
	"github.com/rostislaved/go-clean-architecture/internal/libs/repo-helpers"
)

type BooksRepositoryPostgres struct {
	logger *slog.Logger
	config Config
	DB     *sql.DB
}

func New(l *slog.Logger, cfg Config) *BooksRepositoryPostgres {
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

	return &BooksRepositoryPostgres{
		logger: l,
		config: cfg,
		DB:     db,
	}
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

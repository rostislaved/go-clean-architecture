package books_repository_mongo

import (
	"context"
	"log/slog"
	"time"

	"github.com/rostislaved/go-clean-architecture/internal/libs/helpers"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type BooksRepositoryMongo struct {
	logger *slog.Logger
	config Config
	DB     *mongo.Database
}

func New(l *slog.Logger, cfg Config) *BooksRepositoryMongo {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	credential := options.Credential{
		Username: cfg.User,
		Password: cfg.Password,
	}

	clientOptions := options.Client().ApplyURI(cfg.Host).SetAuth(credential)

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		l.Error(err.Error(), "source", helpers.GetFunctionName())

		panic(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		l.Error(err.Error(), "source", helpers.GetFunctionName())

		panic(err)
	}

	db := client.Database(cfg.Name)

	return &BooksRepositoryMongo{
		logger: l,
		config: cfg,
		DB:     db,
	}
}

type Config struct {
	Name     string
	Host     string `config:"envVar"`
	User     string `config:"envVar"`
	Password string `config:"envVar"`
}

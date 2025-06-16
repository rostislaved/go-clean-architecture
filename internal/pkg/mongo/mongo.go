package mongo

import (
	"context"
	"log/slog"
	"time"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"

	"github.com/rostislaved/go-clean-architecture/internal/pkg/helpers"
)

func New(l *slog.Logger, cfg Config) *mongo.Client {
	credential := options.Credential{
		Username: cfg.User,
		Password: cfg.Password,
	}

	options := options.
		Client().
		ApplyURI(cfg.Host).
		SetAuth(credential)

	client, err := mongo.Connect(options)
	if err != nil {
		l.Error(err.Error(), "source", helpers.GetFunctionName())

		panic(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Ping(ctx, nil)
	if err != nil {
		l.Error(err.Error(), "source", helpers.GetFunctionName())

		panic(err)
	}

	db := client.Database(cfg.Name)

	return client
}

type Config struct {
	Name     string
	Host     string `config:"envVar"`
	User     string `config:"envVar"`
	Password string `config:"envVar"`
}

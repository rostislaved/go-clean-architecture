package employeesRepository

import (
	"context"
	"log/slog"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/rostislaved/go-clean-architecture/internal/pkg/helpers"
)

type DocumentsRepository struct {
	logger *slog.Logger
	config DatabaseMongo
	DB     *mongo.Database
}

func New(l *slog.Logger, cfg DatabaseMongo) *DocumentsRepository {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	credential := options.Credential{
		Username: cfg.User,
		Password: cfg.Password,
	}

	clientOptions := options.Client().ApplyURI(cfg.Host).SetAuth(credential)

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		l.Error(err.Error(), helpers.GetFunctionName())

		panic(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		l.Error(err.Error(), helpers.GetFunctionName())

		panic(err)
	}

	db := client.Database(cfg.Name)

	return &DocumentsRepository{
		logger: l,
		config: cfg,
		DB:     db,
	}
}

type DatabaseMongo struct {
	Name     string
	Host     string `config:"envVar"`
	User     string `config:"envVar"`
	Password string `config:"envVar"`
}

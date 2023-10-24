package employeesRepository

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/rostislaved/go-clean-architecture/internal/app/domain/config"
	"github.com/rostislaved/go-clean-architecture/internal/pkg/helpers"
)

type DocumentsRepository struct {
	logger *slog.Logger
	config config.DatabaseMongo
	DB     *mongo.Database
}

func New(l *slog.Logger, cfg config.DatabaseMongo) *DocumentsRepository {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	credential := options.Credential{
		Username: cfg.User,
		Password: cfg.Password,
	}

	clientOptions := options.Client().ApplyURI(cfg.Host).SetAuth(credential)

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		l.WriteFatal(err.Error(), helpers.GetFunctionName())
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		l.WriteFatal(err.Error(), helpers.GetFunctionName())
	}

	db := client.Database(cfg.Name)

	return &DocumentsRepository{
		logger: l,
		config: cfg,
		DB:     db,
	}
}

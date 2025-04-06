package entity3_repository

import (
	"log/slog"

	"go.mongodb.org/mongo-driver/v2/mongo"
)

type Entity3Repository struct {
	logger *slog.Logger
	config Config
	client *mongo.Client
}

type Config struct{}

func New(l *slog.Logger, cfg Config, client *mongo.Client) *Entity3Repository {
	return &Entity3Repository{
		logger: l,
		config: cfg,
		client: client,
	}
}

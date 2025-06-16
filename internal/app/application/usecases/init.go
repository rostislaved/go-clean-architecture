package usecases

import (
	"context"
	"log/slog"
	"time"

	"github.com/rostislaved/go-clean-architecture/internal/app/domain/entity1"
	"github.com/rostislaved/go-clean-architecture/internal/app/domain/entity5"
)

type UseCases struct {
	logger                *slog.Logger
	config                Config
	entity1Repository     entity1Repository
	gateway               gateway
	kafkaAdapterPublisher Entity1Sender
	natsAdapterPublisher  Entity1Sender
}

type Config struct {
	UpdateInterval time.Duration
}

type entity1Repository interface {
	Get(ctx context.Context, ids []int) (entities []entity1.Entity1, err error)
	Save(ctx context.Context, entities []entity1.Entity1) (createdEntities []entity1.Entity1, err error)
}

type Entity1Sender interface {
	SendEntity1(ctx context.Context, e entity1.Entity1) error
}

type gateway interface {
	Get(ctx context.Context, input struct{}) (entities []entity5.Entity5, err error)
}

func New(
	l *slog.Logger,
	cfg Config,
	repository entity1Repository,
	gateway gateway,
	kafkaAdapterPublisher Entity1Sender,
	natsAdapterPublisher Entity1Sender,
) *UseCases {
	return &UseCases{
		logger:                l,
		config:                cfg,
		entity1Repository:     repository,
		gateway:               gateway,
		kafkaAdapterPublisher: kafkaAdapterPublisher,
		natsAdapterPublisher:  natsAdapterPublisher,
	}
}

package usecases

import (
	"context"
	"log/slog"
	"time"

	"github.com/rostislaved/go-clean-architecture/internal/app/domain/book"
)

type UseCases struct {
	logger                *slog.Logger
	config                Config
	booksRepository       booksRepository
	provider              provider
	kafkaAdapterPublisher bookSender
	natsAdapterPublisher  bookSender
}

type Config struct {
	UpdateInterval time.Duration
}

type booksRepository interface {
	Get(ctx context.Context, ids []int) (books []book.Book, err error)
	Save(ctx context.Context, books []book.Book) (ids []int, err error)
}

type bookSender interface {
	SendBook(ctx context.Context, b book.Book) error
}

type provider interface{}

func New(
	l *slog.Logger,
	cfg Config,
	repository booksRepository,
	provider provider,
	kafkaAdapterPublisher bookSender,
	natsAdapterPublisher bookSender,
) *UseCases {
	return &UseCases{
		logger:                l,
		config:                cfg,
		booksRepository:       repository,
		provider:              provider,
		kafkaAdapterPublisher: kafkaAdapterPublisher,
		natsAdapterPublisher:  natsAdapterPublisher,
	}
}

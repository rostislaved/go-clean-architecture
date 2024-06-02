package kafkaAdapterPublisher

import (
	"context"
	"encoding/json"

	"github.com/segmentio/kafka-go"

	"github.com/rostislaved/go-clean-architecture/internal/app/domain/book"
)

func (a *KafkaAdapterPublisher) SendBook(ctx context.Context, b book.Book) error {
	bookJSONBytes, err := json.Marshal(b)
	if err != nil {
		return err
	}

	message := kafka.Message{
		Key:   []byte("Key"),
		Value: bookJSONBytes,
	}

	err = a.writer.WriteMessages(ctx, message)
	if err != nil {
		return err
	}

	return nil
}

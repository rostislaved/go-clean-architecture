package kafka_adapter_publisher

import (
	"context"
	"encoding/json"
	"time"

	"github.com/rostislaved/go-clean-architecture/internal/app/domain/book"
	"github.com/segmentio/kafka-go"
)

func (a *KafkaAdapterPublisher) SendBook(ctx context.Context, b book.Book) error {
	r := Request(b)

	bookJSONBytes, err := json.Marshal(r)
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

type Request struct {
	ID            int64     `json:"id"`
	Name          string    `json:"name"`
	Author        string    `json:"author"`
	Date          time.Time `json:"date"`
	NumberOfPages int       `json:"number_of_pages"`
}

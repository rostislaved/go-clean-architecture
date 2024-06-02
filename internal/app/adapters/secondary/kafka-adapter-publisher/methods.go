package kafkaAdapterPublisher

import (
	"context"
	"encoding/json"

	"github.com/segmentio/kafka-go"
)

func (a *KafkaAdapterPublisher) SendOrder(ctx context.Context, o book.Book) error {
	orderJSONBytes, err := json.Marshal(o)
	if err != nil {
		return err
	}

	message := kafka.Message{
		Key:   []byte("Key"),
		Value: orderJSONBytes,
	}

	err = a.writer.WriteMessages(ctx, message)
	if err != nil {
		return err
	}

	return nil
}

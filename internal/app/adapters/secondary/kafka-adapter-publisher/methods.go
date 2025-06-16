package kafka_adapter_publisher

import (
	"context"
	"encoding/json"
	"time"

	"github.com/segmentio/kafka-go"

	"github.com/rostislaved/go-clean-architecture/internal/app/domain/entity1"
)

func (a *KafkaAdapterPublisher) SendEntity1(ctx context.Context, entities entity1.Entity1) error {
	request := Request(entities)

	JSONBytes, err := json.Marshal(request)
	if err != nil {
		return err
	}

	message := kafka.Message{
		Key:   []byte("Key"),
		Value: JSONBytes,
	}

	err = a.writer.WriteMessages(ctx, message)
	if err != nil {
		return err
	}

	return nil
}

type Request struct {
	ID     int64     `json:"id"`
	Field1 string    `json:"field1"`
	Field2 int       `json:"field2"`
	Field3 time.Time `json:"field3"`
}

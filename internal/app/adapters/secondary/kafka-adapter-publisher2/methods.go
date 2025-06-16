package kafka_adapter_publisher

import (
	"context"
	"time"

	"github.com/rostislaved/go-clean-architecture/internal/app/domain/entity1"
)

func (a *KafkaAdapterPublisher) SendEntities(ctx context.Context, entities1 ...entity1.Entity1) error {
	return nil
}

type Request struct {
	ID     int64     `json:"id"`
	Field1 string    `json:"field1"`
	Field2 int       `json:"field2"`
	Field3 time.Time `json:"field3"`
}

type Message struct {
	Payload []byte `json:"payload"`
}

package nats_adapter_publisher

import (
	"context"
	"encoding/json"
	"time"

	"github.com/rostislaved/go-clean-architecture/internal/app/domain/entity1"
)

func (a *NatsAdapterPublisher) SendEntity1(ctx context.Context, entity entity1.Entity1) error {
	request := Request(entity)

	JSONBytes, err := json.Marshal(request)
	if err != nil {
		return err
	}

	err = a.publisher.Publish(a.config.Publisher1.Channel, JSONBytes)
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

package kafka_handlers

import (
	"context"
	"encoding/json"
	"time"

	"github.com/rostislaved/go-clean-architecture/internal/app/domain/entity1"
)

func (h KafkaHandlers) SaveEntities1(ctx context.Context, message []byte) (err error) {
	var request Request

	err = json.Unmarshal(message, &request)
	if err != nil {
		return
	}

	entities := request.ToEntity()

	value, err := h.service.Save(ctx, entities)
	if err != nil {
		return
	}

	_ = value

	return
}

type Request struct {
	RequestEntities1 []RequestEntity `json:"entities1"`
}

type RequestEntity struct {
	ID     int64     `json:"id"`
	Field1 string    `json:"field1"`
	Field2 int       `json:"field2"`
	Field3 time.Time `json:"field3"`
}

func (r Request) ToEntity() []entity1.Entity1 {
	entities := make([]entity1.Entity1, 0, len(r.RequestEntities1))

	for _, requestEntity := range r.RequestEntities1 {
		entities = append(entities, entity1.Entity1(requestEntity))
	}

	return entities
}

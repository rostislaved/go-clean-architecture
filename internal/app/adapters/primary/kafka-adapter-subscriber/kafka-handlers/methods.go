package kafka_handlers

import (
	"context"
	"encoding/json"
	"time"

	"github.com/rostislaved/go-clean-architecture/internal/app/domain/book"
)

func (ctr KafkaHandlers) SaveBooks(ctx context.Context, message []byte) (err error) {
	var request Request

	err = json.Unmarshal(message, &request)
	if err != nil {
		return
	}

	books := request.ToEntity()

	value, err := ctr.service.SaveBooks(ctx, books)
	if err != nil {
		return
	}

	_ = value

	return
}

type RequestBook struct {
	ID            int64     `json:"id"`
	Name          string    `json:"name"`
	Author        string    `json:"author"`
	Date          time.Time `json:"date"`
	NumberOfPages int       `json:"number_of_pages"`
}

type Request struct {
	RequestBooks []RequestBook `json:"books"`
}

func (r Request) ToEntity() []book.Book {
	books := make([]book.Book, 0, len(r.RequestBooks))

	for _, requestBook := range r.RequestBooks {
		books = append(books, book.Book(requestBook))
	}

	return books
}

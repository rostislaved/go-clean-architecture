package nats_handlers

import (
	"context"
	"encoding/json"
	"time"

	"github.com/rostislaved/go-clean-architecture/internal/app/domain/book"
)

func (ctr NatsHandlers) SaveBooks(message []byte) (err error) {
	ctx := context.TODO()

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

type Request struct {
	RequestBooks []RequestBook `json:"books"`
}

type RequestBook struct {
	ID            int64     `json:"id"`
	Name          string    `json:"name"`
	Author        string    `json:"author"`
	Date          time.Time `json:"date"`
	NumberOfPages int       `json:"number_of_pages"`
}

func (r Request) ToEntity() []book.Book {
	books := make([]book.Book, 0, len(r.RequestBooks))

	for _, requestBook := range r.RequestBooks {
		books = append(books, book.Book(requestBook))
	}

	return books
}

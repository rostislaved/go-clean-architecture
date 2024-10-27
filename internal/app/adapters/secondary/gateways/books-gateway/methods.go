package books_gateway

import (
	"context"
	"net/url"
	"time"

	"github.com/rostislaved/go-clean-architecture/internal/app/domain/book"
	providerhelpers "github.com/rostislaved/go-clean-architecture/internal/libs/provider-helpers"
)

func (prv *BooksGateway) GetBooks(ctx context.Context, input struct{}) (books []book.Book, err error) {
	ctx, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	values := url.Values{
		"param1": []string{"value1", "value2"},
		"param2": []string{"value1"},
	}

	req := providerhelpers.CreateRequest(ctx, prv.client, prv.config.Endpoints.SignFile)

	var request Request

	req.
		SetQueryParamsFromValues(values).
		SetBody(input).
		ForceContentType("application/json").
		SetResult(&request)

	resp, err := req.Send()
	if err != nil {
		return
	}

	err = providerhelpers.ValidateStatusCode(resp.StatusCode(), resp.Body())
	if err != nil {
		return
	}

	books = request.ToEntity()

	return books, nil
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

package service

import (
	"context"
	"time"

	"github.com/rostislaved/go-clean-architecture/internal/app/domain/book"
)

func (svc *ApiService) Start() {
	ticker := time.NewTicker(svc.config.UpdateInterval)

	for ; true; <-ticker.C {
		err := svc.loop()
		if err != nil {
		}
	}
}

func (svc *ApiService) loop() (err error) {
	return nil
}

func (svc *ApiService) GetBooksByIDs(ctx context.Context, ids []int) (books []book.Order, err error) {
	books, err = svc.repository.GetBooks(ctx, ids)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (svc *ApiService) SaveBooks(ctx context.Context, books []book.Order) (ids []int, err error) {
	ids, err = svc.repository.SaveBooks(ctx, books)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

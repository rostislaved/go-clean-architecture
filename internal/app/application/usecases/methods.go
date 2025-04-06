package usecases

import (
	"context"
	"errors"

	"github.com/rostislaved/go-clean-architecture/internal/app/domain/entity1"
)

var ErrNotFound = errors.New("not found")

func (svc *UseCases) GetBooksByIDs(ctx context.Context, ids []int) (books []entity1.Entity1, err error) {
	books, err = svc.booksRepository.Get(ctx, ids)
	if err != nil {
		return nil, err
	}

	return books, nil
}

func (svc *UseCases) SaveBooks(ctx context.Context, books []entity1.Entity1) (ids []int, err error) {
	ids, err = svc.booksRepository.Save(ctx, books)
	if err != nil {
		return nil, err
	}

	return ids, nil
}

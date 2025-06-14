package usecases

import (
	"context"
	"errors"

	"github.com/rostislaved/go-clean-architecture/internal/app/domain/entity1"
)

var ErrNotFound = errors.New("not found")

func (svc *UseCases) Get(ctx context.Context, ids []int) (entities []entity1.Entity1, err error) {
	entities, err = svc.entity1Repository.Get(ctx, ids)
	if err != nil {
		return nil, err
	}

	return entities, nil
}

func (svc *UseCases) Save(ctx context.Context, entities []entity1.Entity1) (ids []int, err error) {
	ids, err = svc.entity1Repository.Save(ctx, entities)
	if err != nil {
		return nil, err
	}

	return ids, nil
}

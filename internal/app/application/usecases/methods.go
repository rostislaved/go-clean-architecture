package usecases

import (
	"context"
	"errors"

	"github.com/rostislaved/go-clean-architecture/internal/app/domain/entity1"
)

var ErrNotFound = errors.New("not found")

func (uc *UseCases) Get(ctx context.Context, ids []int) (entities []entity1.Entity1, err error) {
	entities, err = uc.entity1Repository.Get(ctx, ids)
	if err != nil {
		return nil, err
	}

	return entities, nil
}

func (uc *UseCases) Save(ctx context.Context, entities []entity1.Entity1) (ids []int64, err error) {
	createdEntities, err := uc.entity1Repository.Save(ctx, entities)
	if err != nil {
		return nil, err
	}

	ids = make([]int64, 0, len(createdEntities))

	for _, entity := range createdEntities {
		ids = append(ids, entity.ID)
	}

	return ids, nil
}

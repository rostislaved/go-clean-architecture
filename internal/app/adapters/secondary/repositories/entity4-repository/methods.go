package entity4_repository

import (
	"github.com/rostislaved/go-clean-architecture/internal/app/domain/entity4"
)

func (repo *Entity4Repository) Get(ids []int64) (entities []entity4.Entity4, err error) {
	//
	// ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	// defer cancel()

	// s := repo.db.Stats()

	return entities, nil
}

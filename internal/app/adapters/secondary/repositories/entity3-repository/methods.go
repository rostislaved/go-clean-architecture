package entity3_repository

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"

	"github.com/rostislaved/go-clean-architecture/internal/app/domain/entity3"
)

func (repo *Entity3Repository) Get(ids []int64) ([]entity3.Entity3, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	doc := bson.D{
		{
			Key: "id",
			Value: bson.D{
				{
					Key:   "$in",
					Value: ids,
				},
			},
		},
	}

	collection := repo.client.Database("database").Collection("entity3_collection")

	cursor, err := collection.Find(ctx, doc)
	if err != nil {
		return nil, err
	}

	defer func() {
		errC := cursor.Close(ctx)
		if errC != nil {
			repo.logger.Info(errC.Error())
		}
	}()

	var entities []entity3.Entity3

	err = cursor.All(ctx, &entities)
	if err != nil {
		return nil, err
	}

	return entities, nil
}

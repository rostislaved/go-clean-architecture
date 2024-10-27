package books_repository_mongo

import (
	"context"
	"log"
	"time"

	"github.com/rostislaved/go-clean-architecture/internal/app/domain/book"
	"go.mongodb.org/mongo-driver/bson"
)

func (repo *BooksRepositoryMongo) Get(ids []int64) ([]book.Book, error) {
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

	collection := repo.DB.Collection("book_collection")

	cursor, err := collection.Find(ctx, doc)
	if err != nil {
		return nil, err
	}

	defer func() {
		errC := cursor.Close(ctx)
		if errC != nil {
			log.Println(errC)
		}
	}()

	var books []book.Book

	err = cursor.All(ctx, &books)
	if err != nil {
		return nil, err
	}

	return books, nil
}

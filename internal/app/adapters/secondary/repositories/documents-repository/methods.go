package employeesRepository

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

func (repo *DocumentsRepository) GetEmployees(ids []int64) ([]book.Employee, error) {
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

	employeeCollection := repo.DB.Collection("employeeCollection")

	cursor, err := employeeCollection.Find(ctx, doc)
	if err != nil {
		return nil, err
	}

	defer func() {
		errC := cursor.Close(ctx)
		if errC != nil {
			log.Println(errC)
		}
	}()

	var employeeList []book.Employee

	err = cursor.All(ctx, &employeeList)
	if err != nil {
		return nil, err
	}

	return employeeList, nil
}

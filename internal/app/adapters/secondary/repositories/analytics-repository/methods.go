package torgRepository

import (
	"context"
	"time"
)

func (repo *AnalyticsRepository) GetSomething(ids []int64) (result []int64, err error) {
	queryString := repo.getGetSomethingQuery(ids)

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	rows, err := repo.DB.QueryContext(ctx, queryString)
	if err != nil {
		return
	}

	result, err = scanGetSomething(rows)
	if err != nil {
		return
	}

	return result, nil
}

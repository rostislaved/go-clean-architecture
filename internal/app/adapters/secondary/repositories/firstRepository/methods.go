package firstRepository

import (
	"context"
	"time"

	"github.com/rostislaved/go-clean-architecture/internal/app/domain/order"
)

func (repo *FirstRepository) Method1() (someStructs []order.Order, err error) {
	queryString := repo.getMethod1Query()

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	rows, err := repo.DB.QueryContext(ctx, queryString)
	if err != nil {
		return
	}

	someStructs, err = scanMethod1(rows)
	if err != nil {
		return
	}

	return
}

package books_repository_clickhouse

import (
	"context"
	"database/sql"
	"time"

	"github.com/rostislaved/go-clean-architecture/internal/app/domain/book"
)

func (repo *BooksRepositoryClickhouse) Get(ids []int64) (books []book.Book, err error) {
	queryString := repo.getGetSomethingQuery(ids)

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	rows, err := repo.DB.QueryContext(ctx, queryString)
	if err != nil {
		return
	}

	books, err = scanGetSomething(rows)
	if err != nil {
		return
	}

	return books, nil
}

func scanGetSomething(rows *sql.Rows) ([]book.Book, error) {
	return nil, nil
}

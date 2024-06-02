package booksRepository

import (
	"context"
	"database/sql"
	"errors"

	sq "github.com/Masterminds/squirrel"

	"github.com/rostislaved/go-clean-architecture/internal/app/domain/book"
)

func (repo *FirstRepository) Get(ctx context.Context, ids []int) (books []book.Book, err error) {
	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

	builder := psql.
		Select(
			"name",
			"author",
			"price",
		).
		From("services.bu_entry_get()").
		Where(sq.Eq{"id": ids})

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	err = repo.DB.SelectContext(ctx, &books, query, args...)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, synchelper.ErrNoData
		}

		return nil, err
	}

	return
}

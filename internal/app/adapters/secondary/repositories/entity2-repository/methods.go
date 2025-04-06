package entity2_repository

import (
	"context"
	"errors"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"

	"github.com/rostislaved/go-clean-architecture/internal/app/application/usecases"
	"github.com/rostislaved/go-clean-architecture/internal/app/domain/entity2"
)

type Model struct{}

func (repo *Entity2Repository) Get(ctx context.Context, ids []int) (entities []entity2.Entity2, err error) {
	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

	builder := psql.
		Select(
			"field1",
			"field2",
			"field3",
		).
		From("entity2").
		Where(sq.Eq{"id": ids})

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := repo.db.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	models, err := pgx.CollectRows(rows, pgx.RowToStructByName[Model])
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) || len(models) == 0 {
			return nil, usecases.ErrNotFound
		}

		return nil, err
	}

	entities = make([]entity2.Entity2, 0, len(models))

	for _, model := range models {
		entities = append(entities, toEntity(model))
	}

	return
}

func toEntity(m Model) (entity entity2.Entity2) {
	// mapping

	return entity
}

func (repo *Entity2Repository) Save(ctx context.Context, entities []entity2.Entity2) (ids []int, err error) {
	return
}

package entity1_repository

import (
	"context"
	"errors"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"

	"github.com/rostislaved/go-clean-architecture/internal/app/application/usecases"
	"github.com/rostislaved/go-clean-architecture/internal/app/domain/entity1"
)

type Model struct{}

func (repo *Entity1Repository) Get(ctx context.Context, ids []int) (entities []entity1.Entity1, err error) {
	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

	builder := psql.
		Select(
			"field1",
			"field2",
			"field3",
		).
		From("entity1").
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

	entities = make([]entity1.Entity1, 0, len(models))

	for _, model := range models {
		entities = append(entities, toEntity(model))
	}

	return entities, nil
}

func toEntity(m Model) (entity entity1.Entity1) {
	// mapping

	return entity
}

func (repo *Entity1Repository) Save(ctx context.Context, entities []entity1.Entity1) (ids []int, err error) {
	return
}

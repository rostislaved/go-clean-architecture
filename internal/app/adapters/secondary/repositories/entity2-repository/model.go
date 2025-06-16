package entity2_repository

import (
	"database/sql"

	"github.com/rostislaved/go-clean-architecture/internal/app/domain/entity2"
)

type model struct {
	ID     sql.NullInt64
	Field1 sql.NullString
	Field2 sql.NullInt64
	Field3 sql.NullTime
}

func (m *model) toEntity() (entity2.Entity2, error) {
	// add fields validation if necessary
	return entity2.Entity2{
		ID:     m.ID.Int64,
		Field1: m.Field1.String,
		Field2: int(m.Field2.Int64),
		Field3: m.Field3.Time,
	}, nil
}

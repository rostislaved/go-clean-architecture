package entity1_repository

import (
	"database/sql"

	"github.com/rostislaved/go-clean-architecture/internal/app/domain/entity1"
)

type model struct {
	ID     sql.NullInt64
	Field1 sql.NullString
	Field2 sql.NullInt64
	Field3 sql.NullTime
}

func (m model) toEntity() (entity1.Entity1, error) {
	// add fields validation if necessary
	return entity1.Entity1{
		ID:     m.ID.Int64,
		Field1: m.Field1.String,
		Field2: int(m.Field2.Int64),
		Field3: m.Field3.Time,
	}, nil
}

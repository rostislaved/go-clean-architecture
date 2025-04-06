package entity2_repository

import (
	"database/sql"

	"github.com/rostislaved/go-clean-architecture/internal/app/domain/entity1"
)

type Entity1DTO struct {
	ID     sql.NullInt64
	Field1 sql.NullString
	Field2 sql.NullInt64
	Field3 sql.NullTime
}

func (dto *Entity1DTO) ToEntity() (entity1.Entity1, error) {
	// add fields validation if necessary
	return entity1.Entity1{
		ID:     dto.ID.Int64,
		Field1: dto.Field1.String,
		Field2: int(dto.Field2.Int64),
		Field3: dto.Field3.Time,
	}, nil
}

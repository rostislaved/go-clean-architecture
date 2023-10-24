package firstRepository

import (
	"database/sql"

	"github.com/rostislaved/go-clean-architecture/internal/app/domain/order"
)

type OrderDTO struct {
	Field1 sql.NullString
	Field2 sql.NullInt64
	Field3 sql.NullBool
}

func (dto *OrderDTO) ToEntity() (order.Order, error) {
	return order.Order{
		Field1: dto.Field1.String,
		Field2: dto.Field2.Int64,
		Field3: dto.Field3.Bool,
	}, nil
}

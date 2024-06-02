package booksRepository

import (
	"database/sql"
)

type OrderDTO struct {
	Field1 sql.NullString
	Field2 sql.NullInt64
	Field3 sql.NullBool
}

func (dto *OrderDTO) ToEntity() (book.Book, error) {
	return book.Book{
		Field1: dto.Field1.String,
		Field2: dto.Field2.Int64,
		Field3: dto.Field3.Bool,
	}, nil
}

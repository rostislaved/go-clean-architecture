package booksRepository

import (
	"database/sql"

	"github.com/rostislaved/go-clean-architecture/internal/app/domain/book"
)

type BookDTO struct {
	ID            sql.NullInt64
	Name          sql.NullString
	Author        sql.NullString
	Date          sql.NullTime
	NumberOfPages sql.NullString
}

func (dto *BookDTO) ToEntity() (book.Book, error) {
	// add fields validation if necessary

	return book.Book{
		ID:            dto.ID.Int64,
		Name:          dto.Name.String,
		Author:        dto.Author.String,
		Date:          dto.Date.Time,
		NumberOfPages: 0,
	}, nil
}

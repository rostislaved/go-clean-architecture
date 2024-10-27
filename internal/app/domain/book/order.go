package book

import "time"

type Book struct {
	ID            int64     `db:"id"`
	Name          string    `db:"name"   validate:"required"`
	Author        string    `db:"author" validate:"required"`
	Date          time.Time `db:"date"   validate:"required,datetime"`
	NumberOfPages int       `db:"author" validate:"required"`
}

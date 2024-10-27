package book

import "time"

type Book struct {
	ID            int64
	Name          string
	Author        string
	Date          time.Time
	NumberOfPages int
}

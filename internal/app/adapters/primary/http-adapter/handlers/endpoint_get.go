package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/rostislaved/go-clean-architecture/internal/app/domain/book"
)

func (h Handlers) Get(w http.ResponseWriter, r *http.Request) {
	var ids []int

	err := json.NewDecoder(r.Body).Decode(&ids)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	books, err := h.service.GetBooksByIDs(r.Context(), ids)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	response := ToResponse(books)

	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)

		return
	}
}

type RequestGet struct{}

func ToResponse(books []book.Book) ResponseGet {
	responseGetBooks := make([]ResponseGetBook, 0, len(books))

	for _, book := range books {
		responseGetBooks = append(responseGetBooks, ResponseGetBook(book))
	}

	response := ResponseGet{responseGetBooks}

	return response
}

type ResponseGet struct {
	ResponseGetBooks []ResponseGetBook `json:"books"`
}

type ResponseGetBook struct {
	ID            int64     `json:"id"`
	Name          string    `json:"name"`
	Author        string    `json:"author"`
	Date          time.Time `json:"date"`
	NumberOfPages int       `json:"number_of_pages"`
}

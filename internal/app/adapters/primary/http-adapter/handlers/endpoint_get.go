package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/rostislaved/go-clean-architecture/internal/app/domain/entity1"
)

func (h Handlers) Get(w http.ResponseWriter, r *http.Request) {
	var request RequestGet

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	entities1, err := h.service.Get(r.Context(), request.IDs)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	response := toResponseGet(entities1)

	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)

		return
	}
}

type RequestGet struct {
	IDs []int `json:"ids"`
}

type ResponseGet struct {
	ResponseGetData []ResponseGetData `json:"data"`
}

type ResponseGetData struct {
	ID     int64     `json:"id"`
	Field1 string    `json:"field1"`
	Field2 int       `json:"field2"`
	Field3 time.Time `json:"field3"`
}

func toResponseGet(entities []entity1.Entity1) ResponseGet {
	responseGet := make([]ResponseGetData, 0, len(entities))

	for _, entity := range entities {
		responseGet = append(responseGet, ResponseGetData(entity))
	}

	response := ResponseGet{responseGet}

	return response
}

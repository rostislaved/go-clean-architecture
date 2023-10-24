package apiController

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/rostislaved/go-clean-architecture/internal/app/domain/order"
)

func (ctr Controller) Get(w http.ResponseWriter, r *http.Request) {
	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	var entityQuery []order.Order

	err = json.Unmarshal(bodyBytes, &entityQuery)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	entity, err := ctr.service.Method1(entityQuery)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	entityJSONBytes, err := json.Marshal(entity)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	_, err = w.Write(entityJSONBytes)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)

		return
	}
}

func (ctr Controller) Post(w http.ResponseWriter, r *http.Request) {}

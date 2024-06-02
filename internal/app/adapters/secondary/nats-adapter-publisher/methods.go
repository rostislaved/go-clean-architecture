package natsAdapterPublisher

import (
	"encoding/json"

	"github.com/rostislaved/go-clean-architecture/internal/app/domain/book"
)

func (a *NatsAdapterPublisher) SendBook(b book.Book) error {
	bookJSONBytes, err := json.Marshal(b)
	if err != nil {
		return err
	}

	err = a.publisher.Publish(a.config.Publisher1.Channel, bookJSONBytes)
	if err != nil {
		return err
	}

	return nil
}

package natsAdapterPublisher

import (
	"encoding/json"
)

func (a *NatsAdapterPublisher) SendOrder(o book.Book) error {
	orderJSONBytes, err := json.Marshal(o)
	if err != nil {
		return err
	}

	err = a.publisher.Publish(a.config.Publisher1.Channel, orderJSONBytes)
	if err != nil {
		return err
	}

	return nil
}

package natsAdapterPublisher

import (
	"encoding/json"

	"github.com/rostislaved/go-clean-architecture/internal/app/domain/order"
)

func (a *NatsAdapterPublisher) SendOrder(o order.Order) error {
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

package natsController

import (
	"encoding/json"

	"github.com/rostislaved/go-clean-architecture/internal/app/domain/order"
)

func (ctr NatsController) DoSomething(message []byte) (err error) {
	var smth []order.Order

	err = json.Unmarshal(message, &smth)
	if err != nil {
		return
	}

	value, err := ctr.service.Method1(smth)
	if err != nil {
		return
	}

	_ = value

	return
}

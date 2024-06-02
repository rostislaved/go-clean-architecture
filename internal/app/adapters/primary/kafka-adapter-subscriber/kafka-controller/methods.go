package kafkaController

import (
	"encoding/json"
)

func (ctr KafkaController) DoSomething(message []byte) (err error) {
	var smth []book.Book

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

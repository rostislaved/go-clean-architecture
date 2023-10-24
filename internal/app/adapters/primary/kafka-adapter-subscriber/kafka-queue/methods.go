package kafkaQueue

import (
	"context"
)

func (queue *KafkaQueue) Subscribe(businessLogicFunc func([]byte) error) {
	for {
		err := queue.processOneMessage(businessLogicFunc)
		if err != nil {
		}
	}
}

func (queue *KafkaQueue) processOneMessage(businessLogicFunc func([]byte) error) (err error) {
	ctx := context.Background()

	message, err := queue.kafkaReader.FetchMessage(ctx) // Тут нельзя делать таймаут. Ибо, если очередь пуста, то тут мы тоже блокируемся и по таймауту получим ошибку. А должны ждать следующего сообщения
	if err != nil {
		return err
	}

	err = businessLogicFunc(message.Value)
	if err != nil {
		return err
	}

	err = queue.kafkaReader.CommitMessages(ctx, message)
	if err != nil {
		return err
	}

	return nil
}

package kafka_queue

import (
	"context"
)

func (queue *KafkaQueue) Subscribe(businessLogicFunc func(context.Context, []byte) error) {
	for {
		err := queue.processOneMessage(businessLogicFunc)
		if err != nil {
			continue
		}
	}
}

func (queue *KafkaQueue) processOneMessage(businessLogicFunc func(context.Context, []byte) error) (err error) {
	ctx := context.TODO()

	message, err := queue.kafkaReader.FetchMessage(ctx) // Тут нельзя делать таймаут. Ибо, если очередь пуста, то тут мы тоже блокируемся и по таймауту получим ошибку. А должны ждать следующего сообщения
	if err != nil {
		return err
	}

	err = businessLogicFunc(ctx, message.Value)
	if err != nil {
		return err
	}

	err = queue.kafkaReader.CommitMessages(ctx, message)
	if err != nil {
		return err
	}

	return nil
}

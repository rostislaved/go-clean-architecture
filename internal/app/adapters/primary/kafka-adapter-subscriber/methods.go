package kafka_adapter_subscriber

import "context"

func (a *KafkaAdapter) Start(ctx context.Context) error {
	a.kafkaQueue.Subscribe(a.kafkaHandlers.SaveEntities1)

	return nil
}

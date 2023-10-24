package kafkaAdapterSubscriber

func (a *KafkaAdapter) Start() {
	a.kafkaQueue.Subscribe(a.kafkaController.DoSomething)
}

package config

import "time"

type NatsAdapterPublisher struct {
	Connection Connection

	Publisher1 Publisher
}

type NatsAdapterSubscriber struct {
	Connection Connection

	Subscriber1 Subscriber
}

type Connection struct {
	Host                 string `config:"envVar"`
	ClusterID            string
	ClientID             string
	AllowMultipleClients bool
	User                 string `config:"envVar"`
	Password             string `config:"envVar"`
}

type Publisher struct {
	Channel string
}

type Subscriber struct {
	Channel             string
	QueueGroup          string
	DurableName         string
	MaxInflight         int
	DeliverAllAvailable bool
	AckWaitTimeout      time.Duration
}

type SubscriberEnv struct {
	Channel             string `config:"envVar"`
	QueueGroup          string
	DurableName         string
	MaxInflight         int
	DeliverAllAvailable bool
	AckWaitTimeout      time.Duration
}

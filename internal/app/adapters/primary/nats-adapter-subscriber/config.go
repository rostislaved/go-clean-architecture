package nats_adapter_subscriber

import (
	"time"
)

type Config struct {
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

type Subscriber struct {
	Channel             string
	QueueGroup          string
	DurableName         string
	MaxInflight         int
	DeliverAllAvailable bool
	AckWaitTimeout      time.Duration
}

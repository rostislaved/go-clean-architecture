package nats_adapter_subscriber

import (
	"time"
)

type Config struct {
	Connection Connection

	Subscriber1 Subscriber
}

type Connection struct {
	Host                 string
	ClusterID            string
	ClientID             string
	AllowMultipleClients bool
	User                 string
	Password             string
}

type Subscriber struct {
	Channel             string
	QueueGroup          string
	DurableName         string
	MaxInflight         int
	DeliverAllAvailable bool
	AckWaitTimeout      time.Duration
}

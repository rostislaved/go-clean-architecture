package natsAdapterSubscriber

import (
	"time"

	"github.com/rostislaved/go-clean-architecture/internal/app/config"
)

type NatsAdapterSubscriberConfig struct {
	Connection config.Connection

	Subscriber1 Subscriber
}

type Subscriber struct {
	Channel             string
	QueueGroup          string
	DurableName         string
	MaxInflight         int
	DeliverAllAvailable bool
	AckWaitTimeout      time.Duration
}

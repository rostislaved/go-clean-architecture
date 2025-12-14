package nats_adapter_publisher

type Config struct {
	Connection Connection

	Publisher1 Publisher
}

type Connection struct {
	Host                 string
	ClusterID            string
	ClientID             string
	AllowMultipleClients bool
	User                 string
	Password             string
}

type Publisher struct {
	Channel string
}

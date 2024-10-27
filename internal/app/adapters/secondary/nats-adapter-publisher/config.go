package nats_adapter_publisher

type Config struct {
	Connection Connection

	Publisher1 Publisher
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

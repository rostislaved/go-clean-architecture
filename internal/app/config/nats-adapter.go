package config

type Connection struct {
	Host                 string `config:"envVar"`
	ClusterID            string
	ClientID             string
	AllowMultipleClients bool
	User                 string `config:"envVar"`
	Password             string `config:"envVar"`
}

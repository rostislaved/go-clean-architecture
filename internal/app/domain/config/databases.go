package config

type DatabaseMongo struct {
	Name     string
	Host     string `config:"envVar"`
	User     string `config:"envVar"`
	Password string `config:"envVar"`
}

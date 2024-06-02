package config

type DatabaseRelational struct {
	Type       string
	Host       string `config:"envVar"`
	Port       string `config:"envVar"`
	User       string `config:"envVar"`
	Password   string `config:"envVar"`
	Name       string
	Procedures map[string]string
}

type DatabaseMongo struct {
	Name     string
	Host     string `config:"envVar"`
	User     string `config:"envVar"`
	Password string `config:"envVar"`
}

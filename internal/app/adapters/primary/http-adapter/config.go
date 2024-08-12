package httpAdapter

import "time"

type HttpAdapterConfig struct {
	Server ServerConfig
	Router RouterConfig
}

type ServerConfig struct {
	Port string
}

type RouterConfig struct {
	Shutdown             shutdown
	Timeout              timeout
	AuthenticationConfig string `config:"envVar"`
	AuthorizationConfig  string `config:"envVar"`
}

type shutdown struct {
	Duration time.Duration
}

type timeout struct {
	Duration time.Duration
}

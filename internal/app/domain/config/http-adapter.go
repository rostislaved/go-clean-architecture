package config

import "time"

type HttpAdapter struct {
	Server Server
	Router Router
}

type Server struct {
	Port string
}

type Router struct {
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

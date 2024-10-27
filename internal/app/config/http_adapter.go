package config

import (
	"time"

	http_server "github.com/rostislaved/go-clean-architecture/internal/libs/http-server"
)

type HttpAdapter struct {
	Server http_server.Config
	Router Router
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

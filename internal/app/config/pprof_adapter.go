package config

import http_server "github.com/rostislaved/go-clean-architecture/internal/libs/http-server"

type PprofAdapter struct {
	Server http_server.Config
}

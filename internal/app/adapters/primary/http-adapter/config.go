package http_adapter

import (
	"github.com/rostislaved/go-clean-architecture/internal/app/adapters/primary/http-adapter/router"
	http_server "github.com/rostislaved/go-clean-architecture/internal/pkg/http-server"
)

type Config struct {
	Server http_server.Config
	Router router.Config
}

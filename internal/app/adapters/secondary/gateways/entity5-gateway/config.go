package entity5_gateway

import providerhelpers "github.com/rostislaved/go-clean-architecture/internal/pkg/provider-helpers"

type Config struct {
	Host      string
	Endpoints Endpoints
}

type Endpoints struct {
	Get    providerhelpers.Endpoint
	Create providerhelpers.Endpoint
}

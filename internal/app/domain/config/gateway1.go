package config

import providerhelpers "github.com/rostislaved/go-clean-architecture/internal/pkg/provider-helpers"

type Gateway1 struct {
	Host      string
	Endpoints Gateway1Endpoints
}

type Gateway1Endpoints struct {
	SignFile       providerhelpers.Endpoint
	GetCertificate providerhelpers.Endpoint
}

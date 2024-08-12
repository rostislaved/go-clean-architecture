package gateway1

import providerhelpers "github.com/rostislaved/go-clean-architecture/internal/pkg/provider-helpers"

type Config struct {
	Host      string
	Endpoints Endpoints
}

type Endpoints struct {
	SignFile       providerhelpers.Endpoint
	GetCertificate providerhelpers.Endpoint
}

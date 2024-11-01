package books_gateway

import providerhelpers "github.com/rostislaved/go-clean-architecture/internal/libs/provider-helpers"

type Config struct {
	Host      string
	Endpoints Endpoints
}

type Endpoints struct {
	SignFile       providerhelpers.Endpoint
	GetCertificate providerhelpers.Endpoint
}

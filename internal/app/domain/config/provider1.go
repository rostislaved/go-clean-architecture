package config

import providerhelpers "github.com/rostislaved/go-clean-architecture/internal/pkg/provider-helpers"

type Gateway1 struct {
	Host      string
	Endpoints Gateway1Endpoints
	AuthHMAC  AuthHMAC
}

type Gateway1Endpoints struct {
	SignFile       providerhelpers.Endpoint
	GetCertificate providerhelpers.Endpoint
}

type AuthHMAC struct {
	ClientID     string `config:"envVar"`
	ClientSecret string `config:"envVar"`
}

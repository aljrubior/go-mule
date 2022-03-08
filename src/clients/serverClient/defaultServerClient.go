package serverClient

import (
	"github.com/aljrubior/standalone-runtime/clients"
	"github.com/aljrubior/standalone-runtime/conf"
)

func NewDefaultServerClient(config conf.ServerClientConfig) DefaultServerClient {
	return DefaultServerClient{
		config: config,
	}
}

type DefaultServerClient struct {
	clients.BaseHttpClient
	config conf.ServerClientConfig
}

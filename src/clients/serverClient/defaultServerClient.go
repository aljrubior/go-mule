package serverClient

import (
	"github.com/aljrubior/go-mule/clients"
	"github.com/aljrubior/go-mule/conf"
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

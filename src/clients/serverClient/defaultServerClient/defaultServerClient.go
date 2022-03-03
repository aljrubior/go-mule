package defaultServerClient

import (
	"github.com/aljrubior/standalone-runtime/clients"
	"github.com/aljrubior/standalone-runtime/conf"
)

type DefaultServerClient struct {
	clients.BaseHttpClient
	config conf.ServerClientConfig
}

//go:build wireinject
// +build wireinject

package wires

import (
	"github.com/aljrubior/go-mule/clients/serverClient"
	"github.com/aljrubior/go-mule/conf"
	"github.com/aljrubior/go-mule/managers/serverManager"
	"github.com/aljrubior/go-mule/services"
	"github.com/google/wire"
)

func InitializeServerManager(config conf.ServerClientConfig) serverManager.ServerManager {

	wire.Build(
		serverClient.NewDefaultServerClient,
		services.NewDefaultServerService,
		serverManager.NewDefaultServerManager,
		wire.Bind(new(serverClient.ServerClient), new(serverClient.DefaultServerClient)),
		wire.Bind(new(services.ServerService), new(services.DefaultServerService)),
		wire.Bind(new(serverManager.ServerManager), new(serverManager.DefaultServerManager)),
	)

	return serverManager.DefaultServerManager{}
}

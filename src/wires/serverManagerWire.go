//go:build wireinject
// +build wireinject

package wires

import (
	"github.com/aljrubior/standalone-runtime/clients/serverClient"
	"github.com/aljrubior/standalone-runtime/conf"
	"github.com/aljrubior/standalone-runtime/managers/serverManager"
	"github.com/aljrubior/standalone-runtime/services"
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

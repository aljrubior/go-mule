//go:build wireinject
// +build wireinject

package wires

import (
	"github.com/aljrubior/go-mule/conf"
	"github.com/aljrubior/go-mule/managers/serverManager"
	"github.com/aljrubior/go-mule/managers/serverRegistrationManager"
	"github.com/google/wire"
)

func InitializeServerRegistrationManager(csrConfig conf.CSRConfig, serverManager serverManager.ServerManager) serverRegistrationManager.ServerRegistrationManager {

	wire.Build(
		serverRegistrationManager.NewDefaultServerRegistrationManager,
		wire.Bind(new(serverRegistrationManager.ServerRegistrationManager), new(serverRegistrationManager.DefaultServerRegistrationManager)),
	)

	return serverRegistrationManager.DefaultServerRegistrationManager{}
}

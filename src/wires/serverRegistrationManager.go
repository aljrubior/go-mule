//go:build wireinject
// +build wireinject

package wires

import (
	"github.com/aljrubior/standalone-runtime/conf"
	"github.com/aljrubior/standalone-runtime/managers/serverManager"
	"github.com/aljrubior/standalone-runtime/managers/serverRegistrationManager"
	"github.com/google/wire"
)

func InitializeServerRegistrationManager(csrConfig conf.CSRConfig, serverManager serverManager.ServerManager) serverRegistrationManager.ServerRegistrationManager {

	wire.Build(
		serverRegistrationManager.NewDefaultServerRegistrationManager,
		wire.Bind(new(serverRegistrationManager.ServerRegistrationManager), new(serverRegistrationManager.DefaultServerRegistrationManager)),
	)

	return serverRegistrationManager.DefaultServerRegistrationManager{}
}

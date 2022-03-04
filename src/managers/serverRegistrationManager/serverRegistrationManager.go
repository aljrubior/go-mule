package serverRegistrationManager

import "github.com/aljrubior/standalone-runtime/managers/serverRegistrationManager/entities"

type ServerRegistrationManager interface {
	Register(token, serverName, muleVersion, agentVersion, environment string) (*entities.ServerEntity, error)
}

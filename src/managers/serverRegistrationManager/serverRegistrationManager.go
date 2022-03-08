package serverRegistrationManager

import "github.com/aljrubior/go-mule/managers/serverRegistrationManager/entities"

type ServerRegistrationManager interface {
	Register(token, serverName, muleVersion, agentVersion, environment string) (*entities.ServerEntity, error)
}

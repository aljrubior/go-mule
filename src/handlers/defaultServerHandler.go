package handlers

import "github.com/aljrubior/standalone-runtime/managers/serverRegistrationManager"

const (
	MULE_VERSION  = "4.4.0"
	AGENT_VERSION = "2.4.27"
	ENVIRONMENT   = "qax"
)

func NewDefaultServerHandler(serverRegistrationManager serverRegistrationManager.ServerRegistrationManager) DefaultServerHandler {
	return DefaultServerHandler{
		serverRegistrationManager: serverRegistrationManager,
	}
}

type DefaultServerHandler struct {
	serverRegistrationManager serverRegistrationManager.ServerRegistrationManager
}

func (handler DefaultServerHandler) CreateServer(token, serverName string) error {

	entity, err := handler.serverRegistrationManager.Register(token, serverName, MULE_VERSION, AGENT_VERSION, ENVIRONMENT)

	if err != nil {
		return err
	}

	println(entity.Certificate)

	return nil
}

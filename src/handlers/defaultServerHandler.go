package handlers

import (
	"github.com/aljrubior/standalone-runtime/managers/serverRegistrationManager"
	"github.com/aljrubior/standalone-runtime/writers"
)

func NewDefaultServerHandler(serverRegistrationManager serverRegistrationManager.ServerRegistrationManager) DefaultServerHandler {
	return DefaultServerHandler{
		serverRegistrationManager: serverRegistrationManager,
	}
}

type DefaultServerHandler struct {
	serverRegistrationManager serverRegistrationManager.ServerRegistrationManager
}

func (handler DefaultServerHandler) CreateServer(token, serverName, muleVersion, agentVersion, environment string) error {

	entity, err := handler.serverRegistrationManager.Register(token, serverName, muleVersion, agentVersion, environment)

	if err != nil {
		return err
	}

	err = writers.NewCertificateWriter(entity.PrivateKey, entity.Certificate).WriteFile()

	return nil
}

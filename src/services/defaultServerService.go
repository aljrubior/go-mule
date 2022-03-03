package services

import (
	"github.com/aljrubior/standalone-runtime/clients/serverClient"
	"github.com/aljrubior/standalone-runtime/clients/serverClient/responses"
)

func NewDefaultServerService(serverClient serverClient.ServerClient) DefaultServerService {
	return DefaultServerService{
		serverClient: serverClient,
	}
}

type DefaultServerService struct {
	serverClient serverClient.ServerClient
}

func (serverService DefaultServerService) RegisterServer(token, agentVersion string, body []byte) (*responses.ServerRegistrationResponse, error) {

	return serverService.serverClient.RegisterServer(token, agentVersion, body)
}

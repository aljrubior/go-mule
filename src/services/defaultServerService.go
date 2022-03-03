package services

import (
	"encoding/json"
	"github.com/aljrubior/standalone-runtime/clients/serverClient"
	"github.com/aljrubior/standalone-runtime/clients/serverClient/responses"
	"github.com/aljrubior/standalone-runtime/managers/serverManager/requests"
)

func NewDefaultServerService(serverClient serverClient.ServerClient) DefaultServerService {
	return DefaultServerService{
		serverClient: serverClient,
	}
}

type DefaultServerService struct {
	serverClient serverClient.ServerClient
}

func (serverService DefaultServerService) RegisterServer(token, agentVersion string, request requests.ServerRegistrationRequest) (*responses.ServerRegistrationResponse, error) {

	data, err := json.Marshal(request)

	if err != nil {
		return nil, err
	}

	return serverService.serverClient.RegisterServer(token, agentVersion, data)
}

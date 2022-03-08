package services

import (
	"encoding/json"
	"github.com/aljrubior/go-mule/clients/serverClient"
	"github.com/aljrubior/go-mule/clients/serverClient/responses"
	"github.com/aljrubior/go-mule/managers/serverManager/requests"
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

package serverManager

import (
	"github.com/aljrubior/go-mule/clients/serverClient/responses"
	"github.com/aljrubior/go-mule/managers/serverManager/requests"
	"github.com/aljrubior/go-mule/services"
)

func NewDefaultServerManager(serverService services.ServerService) DefaultServerManager {
	return DefaultServerManager{
		serverService: serverService,
	}
}

type DefaultServerManager struct {
	serverService services.ServerService
}

func (serverManager DefaultServerManager) RegisterServer(token string, request requests.ServerRegistrationRequest) (*responses.ServerRegistrationResponse, error) {

	return serverManager.serverService.RegisterServer(token, request.AgentVersion, request)
}

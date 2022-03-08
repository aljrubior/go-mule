package serverManager

import (
	"github.com/aljrubior/go-mule/clients/serverClient/responses"
	"github.com/aljrubior/go-mule/managers/serverManager/requests"
)

type ServerManager interface {
	RegisterServer(token string, request requests.ServerRegistrationRequest) (*responses.ServerRegistrationResponse, error)
}

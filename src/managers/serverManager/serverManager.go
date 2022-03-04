package serverManager

import (
	"github.com/aljrubior/standalone-runtime/clients/serverClient/responses"
	"github.com/aljrubior/standalone-runtime/managers/serverManager/requests"
)

type ServerManager interface {
	RegisterServer(token string, request requests.ServerRegistrationRequest) (*responses.ServerRegistrationResponse, error)
}

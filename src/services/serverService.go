package services

import (
	"github.com/aljrubior/standalone-runtime/clients/serverClient/responses"
	"github.com/aljrubior/standalone-runtime/managers/serverManager/requests"
)

type ServerService interface {
	RegisterServer(token, agentVersion string, request requests.ServerRegistrationRequest) (*responses.ServerRegistrationResponse, error)
}

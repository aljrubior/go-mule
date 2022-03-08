package services

import (
	"github.com/aljrubior/go-mule/clients/serverClient/responses"
	"github.com/aljrubior/go-mule/managers/serverManager/requests"
)

type ServerService interface {
	RegisterServer(token, agentVersion string, request requests.ServerRegistrationRequest) (*responses.ServerRegistrationResponse, error)
}

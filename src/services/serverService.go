package services

import "github.com/aljrubior/standalone-runtime/clients/serverClient/responses"

type ServerService interface {
	RegisterServer(token, agentVersion string, body []byte) (*responses.ServerRegistrationResponse, error)
}

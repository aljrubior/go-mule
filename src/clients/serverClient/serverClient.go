package serverClient

import "github.com/aljrubior/standalone-runtime/clients/serverClient/responses"

type ServerClient interface {
	RegisterServer(token, agentVersion string, body []byte) (*responses.ServerRegistrationResponse, error)
}

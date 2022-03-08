package serverClient

import "github.com/aljrubior/go-mule/clients/serverClient/responses"

type ServerClient interface {
	RegisterServer(token, agentVersion string, body []byte) (*responses.ServerRegistrationResponse, error)
}

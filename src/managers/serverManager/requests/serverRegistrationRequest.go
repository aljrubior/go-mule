package requests

func NewServerRegistrationRequest(muleVersion, gatewayVersion, muleLabel, csr, agentVersion, environment string) ServerRegistrationRequest {
	return ServerRegistrationRequest{
		MuleVersion:    muleVersion,
		GatewayVersion: gatewayVersion,
		MuleLabel:      muleLabel,
		CSR:            csr,
		AgentVersion:   agentVersion,
		Environment:    environment,
	}
}

type ServerRegistrationRequest struct {
	MuleVersion    string `json:"muleVersion"`
	GatewayVersion string `json:"gatewayVersion"`
	MuleLabel      string `json:"muleLabel"`
	CSR            string `json:"signatureRequest"`
	AgentVersion   string `json:"agentVersion"`
	Environment    string `json:"environment"`
}

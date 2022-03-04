package serverRegistrationManager

import (
	"github.com/aljrubior/standalone-runtime/conf"
	"github.com/aljrubior/standalone-runtime/managers/serverManager"
	"github.com/aljrubior/standalone-runtime/managers/serverManager/requests"
	"github.com/aljrubior/standalone-runtime/managers/serverRegistrationManager/entities"
	"github.com/aljrubior/standalone-runtime/security"
)

const (
	RSA_KEY_SIZE_DEFAULT = 2048
)

func NewServerRegistrationManager(
	serverManager serverManager.ServerManager,
	csrConfig conf.CSRConfig) ServerRegistrationManager {

	return ServerRegistrationManager{
		serverManager: serverManager,
		csrConfig:     csrConfig,
	}
}

type ServerRegistrationManager struct {
	serverManager serverManager.ServerManager
	csrConfig     conf.CSRConfig
}

func (registrationManager ServerRegistrationManager) Register(token, serverName, muleVersion, agentVersion, environment string) (*entities.ServerEntity, error) {

	keyPairGenerator := security.NewRSAKeyPairGenerator(registrationManager.csrConfig, RSA_KEY_SIZE_DEFAULT)

	privateKey, csr, err := keyPairGenerator.GenerateKeyCsr()

	if err != nil {
		return nil, err
	}

	csrAsString := NewCSRWrapper(csr).ToString()

	request := requests.NewServerRegistrationRequest(muleVersion, muleVersion, serverName, csrAsString, agentVersion, environment)

	response, err := registrationManager.serverManager.RegisterServer(token, request)

	if err != nil {
		return nil, err
	}

	cert := NewSignedCertificateWrapper(response.Certificate).Parse()

	return entities.NewServerEntity(
		privateKey,
		[]byte(cert),
		response.EnvironmentUrls.MCMWebsocketUrl,
		response.EnvironmentUrls.MCMWebsocketUrl), nil
}

package serverRegistrationManager

import (
	"github.com/aljrubior/go-mule/conf"
	"github.com/aljrubior/go-mule/managers/serverManager"
	"github.com/aljrubior/go-mule/managers/serverManager/requests"
	"github.com/aljrubior/go-mule/managers/serverRegistrationManager/entities"
	"github.com/aljrubior/go-mule/security"
)

const (
	RSA_KEY_SIZE_DEFAULT = 2048
)

func NewDefaultServerRegistrationManager(
	serverManager serverManager.ServerManager,
	csrConfig conf.CSRConfig) DefaultServerRegistrationManager {

	return DefaultServerRegistrationManager{
		serverManager: serverManager,
		csrConfig:     csrConfig,
	}
}

type DefaultServerRegistrationManager struct {
	serverManager serverManager.ServerManager
	csrConfig     conf.CSRConfig
}

func (registrationManager DefaultServerRegistrationManager) Register(token, serverName, muleVersion, agentVersion, environment string) (*entities.ServerEntity, error) {

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

	caCert := NewSignedCertificateWrapper(response.CACertificate).Parse()

	return entities.NewServerEntity(
		privateKey,
		[]byte(cert),
		[]byte(caCert),
		response.EnvironmentUrls.MCMWebsocketUrl,
		response.EnvironmentUrls.MCMWebsocketUrl), nil
}

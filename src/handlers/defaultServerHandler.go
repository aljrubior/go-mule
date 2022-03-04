package handlers

import (
	"fmt"
	"github.com/aljrubior/standalone-runtime/managers/serverRegistrationManager"
	"github.com/aljrubior/standalone-runtime/runtime"
	"github.com/aljrubior/standalone-runtime/security"
	"github.com/aljrubior/standalone-runtime/writers"
	"io/ioutil"
)

func NewDefaultServerHandler(serverRegistrationManager serverRegistrationManager.ServerRegistrationManager) DefaultServerHandler {
	return DefaultServerHandler{
		serverRegistrationManager: serverRegistrationManager,
	}
}

type DefaultServerHandler struct {
	serverRegistrationManager serverRegistrationManager.ServerRegistrationManager
}

func (handler DefaultServerHandler) CreateServer(token, serverName, muleVersion, agentVersion, environment string) error {

	entity, err := handler.serverRegistrationManager.Register(token, serverName, muleVersion, agentVersion, environment)

	if err != nil {
		return err
	}

	err = writers.NewCertificateWriter(entity.PrivateKey, entity.Certificate, entity.CACertificate).WriteFile()

	return nil
}

func (handler DefaultServerHandler) StartServer(serverId string) error {

	privateKeyPath := fmt.Sprintf("./%s/%s.key", serverId, serverId)
	certificatePath := fmt.Sprintf("./%s/%s.pem", serverId, serverId)
	caCertificatePath := fmt.Sprintf("./%s/ca.pem", serverId)

	certificate, err := ioutil.ReadFile(certificatePath)

	if err != nil {
		return err
	}

	contextId, err := security.NewCertificateWrapper(certificate).GetOrganizationalUnit()

	if err != nil {
		return err
	}

	runtime := runtime.NewStandaloneRuntime(serverId, contextId, certificatePath, privateKeyPath, caCertificatePath)

	runtime.Start()

	return nil
}

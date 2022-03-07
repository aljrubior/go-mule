package handlers

import (
	"fmt"
	"github.com/aljrubior/standalone-runtime/managers/serverRegistrationManager"
	"github.com/aljrubior/standalone-runtime/runtime"
	"github.com/aljrubior/standalone-runtime/security"
	"github.com/aljrubior/standalone-runtime/writers"
	"io/ioutil"
)

const (
	TOTAL_FLOWS_PER_APPLICATION_DEFAULT = 10
)

func NewDefaultServerHandler(serverRegistrationManager serverRegistrationManager.ServerRegistrationManager) DefaultServerHandler {
	return DefaultServerHandler{
		serverRegistrationManager: serverRegistrationManager,
	}
}

type DefaultServerHandler struct {
	serverRegistrationManager serverRegistrationManager.ServerRegistrationManager
}

func (t DefaultServerHandler) CreateServer(token, serverName, muleVersion, agentVersion, environment string) error {

	entity, err := t.serverRegistrationManager.Register(token, serverName, muleVersion, agentVersion, environment)

	if err != nil {
		return err
	}

	err = writers.NewCertificateWriter(entity.PrivateKey, entity.Certificate, entity.CACertificate).WriteFile()

	return nil
}

func (t DefaultServerHandler) StartServer(serverId string, totalFlowsPerApp int) error {

	totalFlows := t.getTotalFlows(totalFlowsPerApp)
	totalFixedSchPerApp, totalCronSchPerApp := t.getSchedulerCountPerApplication(totalFlows)

	println(fmt.Sprintf("Runtime Configuration - Flows Per Application [total: '%v']", totalFlows))
	println(fmt.Sprintf("Runtime Configuration - Schedulers Per Application [fixedFrequency: '%v' cron: '%v']\n", totalFixedSchPerApp, totalCronSchPerApp))

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

	runtime := runtime.NewStandaloneRuntime(serverId, contextId, certificatePath, privateKeyPath, caCertificatePath, totalFixedSchPerApp, totalCronSchPerApp)

	runtime.Start()

	return nil
}

func (t DefaultServerHandler) getTotalFlows(count int) int {

	if count > 0 {
		return count
	}

	return TOTAL_FLOWS_PER_APPLICATION_DEFAULT
}

func (t DefaultServerHandler) getSchedulerCountPerApplication(totalFlowsPerApplication int) (int, int) {

	totalFixedSchedulers := totalFlowsPerApplication / 2
	totalCronSchedulers := totalFlowsPerApplication - totalFixedSchedulers

	return totalFixedSchedulers, totalCronSchedulers
}

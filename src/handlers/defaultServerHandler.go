package handlers

import (
	"crypto/tls"
	"fmt"
	"github.com/aljrubior/standalone-runtime/clients/metricClient"
	"github.com/aljrubior/standalone-runtime/managers/configManager/defaultConfigManager"
	"github.com/aljrubior/standalone-runtime/managers/metricManager"
	"github.com/aljrubior/standalone-runtime/managers/metricManager/requests"
	"github.com/aljrubior/standalone-runtime/managers/serverRegistrationManager"
	"github.com/aljrubior/standalone-runtime/runtime"
	"github.com/aljrubior/standalone-runtime/security"
	"github.com/aljrubior/standalone-runtime/services"
	"github.com/aljrubior/standalone-runtime/writers"
	"io/ioutil"
	"time"
)

const (
	TOTAL_FLOWS_PER_APPLICATION_DEFAULT = 10
)

func NewDefaultServerHandler(
	serverRegistrationManager serverRegistrationManager.ServerRegistrationManager,
	configManager defaultConfigManager.DefaultConfigManager) DefaultServerHandler {
	return DefaultServerHandler{
		serverRegistrationManager: serverRegistrationManager,
		configManager:             configManager,
	}
}

type DefaultServerHandler struct {
	serverRegistrationManager serverRegistrationManager.ServerRegistrationManager
	configManager             defaultConfigManager.DefaultConfigManager
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

	tlsConfig := security.NewTLSConfigBuilder(certificatePath, privateKeyPath, caCertificatePath).Build()

	metricManager := t.buildMetricManager(tlsConfig)

	runtime := runtime.NewStandaloneRuntime(serverId, contextId, tlsConfig, totalFixedSchPerApp, totalCronSchPerApp)

	go t.applicationMetricsSender(&runtime, metricManager)

	go t.serverMetricsSender(metricManager)

	runtime.Start()

	return nil
}

func (t DefaultServerHandler) applicationMetricsSender(runtime *runtime.StandaloneRuntime, metricManager metricManager.MetricManager) {
	for {
		for _, v := range *runtime.GetApplications() {
			metrics := requests.NewApplicationMetricBuilder(v).Build()
			metricManager.PostApplicationMetrics(v.Name, metrics)
		}
		time.Sleep(60 * time.Second)
	}
}

func (t DefaultServerHandler) serverMetricsSender(metricManager metricManager.MetricManager) {
	for {
		now := time.Now()
		metrics := requests.NewServerMetricRequestBuilder(now.Format(time.RFC3339)).Build()

		metricManager.PostServerMetrics(metrics)

		time.Sleep(60 * time.Second)
	}
}

func (t DefaultServerHandler) buildMetricManager(tlsConfig *tls.Config) metricManager.DefaultMetricManager {

	metricClient := metricClient.NewDefaultMetricClient(t.configManager.GetMetricClientConfig(), tlsConfig)
	metricService := services.NewDefaultMetricService(&metricClient)
	return metricManager.NewDefaultMetricManager(&metricService)
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

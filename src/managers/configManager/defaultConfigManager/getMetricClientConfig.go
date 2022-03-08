package defaultConfigManager

import "github.com/aljrubior/go-mule/conf"

func (manager *DefaultConfigManager) GetMetricClientConfig() *conf.MetricClientConfig {

	if manager.metricClientConfig == nil {
		manager.metricClientConfig = &conf.MetricClientConfig{
			Protocol:         manager.mainConfig.Metrics.Protocol,
			Host:             manager.mainConfig.Metrics.Host,
			Port:             manager.mainConfig.Metrics.Port,
			ApplicationsPath: manager.mainConfig.Metrics.ApplicationsPath,
			ServersPath:      manager.mainConfig.Metrics.ServersPath,
		}
	}

	return manager.metricClientConfig
}

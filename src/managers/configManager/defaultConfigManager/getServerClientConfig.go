package defaultConfigManager

import "github.com/aljrubior/go-mule/conf"

func (manager *DefaultConfigManager) GetServerClientConfig() *conf.ServerClientConfig {

	if manager.serverClientConfig == nil {
		manager.serverClientConfig = &conf.ServerClientConfig{
			Protocol:    manager.mainConfig.Anypoint.Protocol,
			Host:        manager.mainConfig.Anypoint.Host,
			Port:        manager.mainConfig.Anypoint.Port,
			ServersPath: manager.mainConfig.Anypoint.Resources.Servers.ServersPath,
		}
	}

	return manager.serverClientConfig
}

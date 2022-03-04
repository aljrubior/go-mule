package defaultConfigManager

import "github.com/aljrubior/standalone-runtime/conf"

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

package defaultConfigManager

import "github.com/aljrubior/standalone-runtime/conf"

func (manager *DefaultConfigManager) GetServerConfigFile() *conf.ServerConfigFile {

	if manager.serverConfigFile == nil {
		manager.serverConfigFile = &conf.ServerConfigFile{
			Protocol:    manager.mainConfig.Anypoint.Protocol,
			Host:        manager.mainConfig.Anypoint.Host,
			Port:        manager.mainConfig.Anypoint.Port,
			ServersPath: manager.mainConfig.Anypoint.Resources.Servers.ServersPath,
		}
	}

	return manager.serverConfigFile
}

package defaultConfigManager

import (
	"github.com/aljrubior/go-mule/conf"
)

type DefaultConfigManager struct {
	configDir,
	configFile string
	mainConfig conf.MainConfig

	csrConfig          *conf.CSRConfig
	serverClientConfig *conf.ServerClientConfig
	metricClientConfig *conf.MetricClientConfig
}

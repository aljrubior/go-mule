package defaultConfigManager

import (
	"github.com/aljrubior/standalone-runtime/conf"
)

type DefaultConfigManager struct {
	configDir,
	configFile string
	mainConfig conf.MainConfig

	csrConfig          *conf.CSRConfig
	serverClientConfig *conf.ServerClientConfig
}

package defaultConfigManager

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

func (manager *DefaultConfigManager) loadConfiguration() error {

	anypointHome, err := manager.getConfigDir()

	if err != nil {
		return err
	}

	filePath := fmt.Sprintf("%s/%s", anypointHome, manager.configFile)

	data, err := ioutil.ReadFile(filePath)

	if err != nil {
		return err
	}

	err = yaml.Unmarshal(data, &manager.mainConfig)
	
	return err
}

package defaultConfigManager

import (
	"fmt"
	"os"
)

func (manager DefaultConfigManager) getConfigDir() (string, error) {
	homeDir, err := os.UserHomeDir()

	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s/%s", homeDir, manager.configDir), nil
}

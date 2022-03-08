package defaultConfigManager

func NewDefaultConfigManager() DefaultConfigManager {
	newManager := DefaultConfigManager{
		configDir:  ".anypoint",
		configFile: "go-mule.yaml",
	}

	newManager.loadConfiguration()

	return newManager
}

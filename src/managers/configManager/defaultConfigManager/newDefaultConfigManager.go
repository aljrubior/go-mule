package defaultConfigManager

func NewDefaultConfigManager() DefaultConfigManager {
	newManager := DefaultConfigManager{
		configDir:  ".anypoint",
		configFile: "standalone-runtime.yaml",
	}

	newManager.loadConfiguration()

	return newManager
}

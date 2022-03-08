package conf

type Metrics struct {
	Protocol         string `yaml:"protocol"`
	Host             string `yaml:"host"`
	Port             string `yaml:"port"`
	ApplicationsPath string `yaml:"applicationsPath"`
	ServersPath      string `yaml:"serversPath"`
}

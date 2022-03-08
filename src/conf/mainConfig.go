package conf

type MainConfig struct {
	Anypoint Anypoint `yaml:"anypoint""`
	Metrics  Metrics  `yaml:"metrics"`
	Security Security `yaml:"security"`
}

package conf

type MainConfig struct {
	Anypoint Anypoint `yaml:"anypoint""`
	Security Security `yaml:"security"`
}

package conf

type Anypoint struct {
	Protocol  string    `yaml:"protocol"`
	Host      string    `yaml:"host"`
	Port      string    `yaml:"port"`
	Resources Resources `yaml:"resources"`
}

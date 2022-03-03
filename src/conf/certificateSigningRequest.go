package conf

type CertificateSigningRequest struct {
	commonName       string `yaml:"commonName"`
	organizationUnit string `yaml:"organizationUnit"`
	organization     string `yaml:"organization"`
	locality         string `yaml:"locality"`
	province         string `yaml:"province"`
	country          string `yaml:"country"`
}

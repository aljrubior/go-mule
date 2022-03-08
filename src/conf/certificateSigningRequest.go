package conf

type CertificateSigningRequest struct {
	CommonName       string `yaml:"commonName"`
	OrganizationUnit string `yaml:"organizationUnit"`
	Organization     string `yaml:"organization"`
	Locality         string `yaml:"locality"`
	Province         string `yaml:"province"`
	Country          string `yaml:"country"`
}

package defaultConfigManager

import "github.com/aljrubior/standalone-runtime/conf"

func (manager *DefaultConfigManager) GetCSRConfig() *conf.CSRConfig {

	if manager.csrConfig == nil {
		manager.csrConfig = &conf.CSRConfig{
			CommonName:       manager.mainConfig.Security.CertificateSigningRequest.CommonName,
			OrganizationUnit: manager.mainConfig.Security.CertificateSigningRequest.OrganizationUnit,
			Organization:     manager.mainConfig.Security.CertificateSigningRequest.Organization,
			Locality:         manager.mainConfig.Security.CertificateSigningRequest.Locality,
			Province:         manager.mainConfig.Security.CertificateSigningRequest.Province,
			Country:          manager.mainConfig.Security.CertificateSigningRequest.Country,
		}
	}

	return manager.csrConfig
}

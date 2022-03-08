package security

import "github.com/aljrubior/go-mule/conf"

func NewRSAKeyPairGenerator(config conf.CSRConfig, keySize int) RSAKeyPairGenerator {

	return RSAKeyPairGenerator{
		config: config,
		bits:   keySize,
	}
}

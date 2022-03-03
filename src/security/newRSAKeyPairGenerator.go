package security

import "github.com/aljrubior/standalone-runtime/conf"

func NewRSAKeyPairGenerator(config conf.CSRConfig, keySize int) RSAKeyPairGenerator {

	return RSAKeyPairGenerator{
		config: config,
		bits:   keySize,
	}
}

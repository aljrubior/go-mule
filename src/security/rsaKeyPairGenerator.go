package security

import "github.com/aljrubior/standalone-runtime/conf"

type RSAKeyPairGenerator struct {
	config conf.CSRConfig
	bits   int
}

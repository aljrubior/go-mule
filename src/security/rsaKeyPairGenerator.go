package security

import "github.com/aljrubior/go-mule/conf"

type RSAKeyPairGenerator struct {
	config conf.CSRConfig
	bits   int
}

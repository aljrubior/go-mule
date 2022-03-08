package security

import (
	"bytes"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
)

func (generator RSAKeyPairGenerator) marshalPrivateKey(privateKey *rsa.PrivateKey) (*bytes.Buffer, error) {

	privateKeyBytes := x509.MarshalPKCS1PrivateKey(privateKey)

	block := &pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: privateKeyBytes,
	}

	var bytesBuffer bytes.Buffer

	err := pem.Encode(&bytesBuffer, block)

	if err != nil {
		return nil, generator.ThrowRSAKeyPairGeneratorError("Error when encode private key pem", err)
	}

	return &bytesBuffer, nil
}

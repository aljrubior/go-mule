package security

import (
	"crypto/rand"
	"crypto/rsa"
)

func (generator RSAKeyPairGenerator) GenerateKeyPair() ([]byte, []byte, error) {

	key, err := rsa.GenerateKey(rand.Reader, generator.bits)

	if err != nil {
		return nil, nil, generator.ThrowRSAKeyPairGeneratorError("Cannot generate RSA key", err)
	}

	privateKey, err := generator.marshalPrivateKey(key)

	if err != nil {
		return nil, nil, err
	}

	csr, err := generator.generateCSR(key)

	return privateKey.Bytes(), csr.Bytes(), nil

}

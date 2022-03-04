package security

import (
	"crypto/x509"
	"encoding/pem"
)

func NewCertificateWrapper(data []byte) CertificateWrapper {
	return CertificateWrapper{
		data: data,
	}
}

type CertificateWrapper struct {
	data []byte
}

func (wrapper CertificateWrapper) GetCommonName() (string, error) {

	block, _ := pem.Decode(wrapper.data)

	cert, err := x509.ParseCertificate(block.Bytes)

	if err != nil {
		return "", err
	}

	return cert.Subject.CommonName, nil
}

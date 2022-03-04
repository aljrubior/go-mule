package security

import "crypto/x509"

func NewCertificateWrapper(data []byte) CertificateWrapper {
	return CertificateWrapper{
		data: data,
	}
}

type CertificateWrapper struct {
	data []byte
}

func (wrapper CertificateWrapper) GetCommonName() (string, error) {

	cert, err := x509.ParseCertificate(wrapper.data)

	if err != nil {
		return "", err
	}

	return cert.Subject.CommonName, nil
}

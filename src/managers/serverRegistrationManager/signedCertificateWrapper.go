package serverRegistrationManager

import "strings"

func NewSignedCertificateWrapper(certificate string) SignedCertificateWrapper {

	return SignedCertificateWrapper{
		certificate: certificate,
	}
}

type SignedCertificateWrapper struct {
	certificate string
}

func (wrapper SignedCertificateWrapper) Parse() string {

	cert := strings.Replace(wrapper.certificate, " REQUEST", "", 2)

	certAsArray := strings.Split(cert, "\\n")

	return strings.Join(certAsArray, "\n")
}

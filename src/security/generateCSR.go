package security

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
)

func (generator RSAKeyPairGenerator) generateCSR(privateKey *rsa.PrivateKey) (*bytes.Buffer, error) {

	subj := pkix.Name{
		CommonName:         generator.config.CommonName,
		OrganizationalUnit: []string{generator.config.OrganizationUnit},
		Organization:       []string{generator.config.Organization},
		Locality:           []string{generator.config.Locality},
		Province:           []string{generator.config.Province},
		Country:            []string{generator.config.Country},
	}

	csr := x509.CertificateRequest{
		Subject:            subj,
		SignatureAlgorithm: x509.SHA256WithRSA,
	}

	csrBytes, err := x509.CreateCertificateRequest(rand.Reader, &csr, privateKey)

	if err != nil {
		return nil, generator.ThrowRSAKeyPairGeneratorError("Error when creating csr", err)
	}

	block := &pem.Block{
		Type:  "CERTIFICATE REQUEST",
		Bytes: csrBytes,
	}

	var bytesBuffer bytes.Buffer

	err = pem.Encode(&bytesBuffer, block)

	if err != nil {
		return nil, generator.ThrowRSAKeyPairGeneratorError("Error when encode csr pem", err)
	}

	return &bytesBuffer, nil
}

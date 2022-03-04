package writers

import (
	"fmt"
	"github.com/aljrubior/standalone-runtime/security"
	"os"
)

func NewCertificateWriter(privateKey, certificate, caCertificate []byte) CertificateWriter {
	return CertificateWriter{
		privateKey:    privateKey,
		certificate:   certificate,
		caCertificate: caCertificate,
	}
}

type CertificateWriter struct {
	privateKey,
	certificate,
	caCertificate []byte
}

func (writer CertificateWriter) WriteFile() error {

	commonName, err := security.NewCertificateWrapper(writer.certificate).GetCommonName()

	if err != nil {
		return err
	}

	workspaceDir := fmt.Sprintf("./%s", commonName)

	if err := os.MkdirAll(workspaceDir, 0755); err != nil {
		println(err.Error())
		return err
	}

	pemFileName := fmt.Sprintf("%s/%s.pem", workspaceDir, commonName)
	keyFileName := fmt.Sprintf("%s/%s.key", workspaceDir, commonName)
	caFileName := fmt.Sprintf("%s/ca.pem", workspaceDir)

	pemFile, err := os.Create(pemFileName)

	if err != nil {
		return err
	}

	keyFile, err := os.Create(keyFileName)

	if err != nil {
		return err
	}

	caFile, err := os.Create(caFileName)

	if err != nil {
		return err
	}

	_, err = pemFile.Write(writer.certificate)

	if err != nil {
		return err
	}

	_, err = keyFile.Write(writer.privateKey)

	if err != nil {
		return err
	}

	_, err = caFile.Write(writer.caCertificate)

	if err != nil {
		return err
	}

	println(fmt.Sprintf("Standalone created successfully [Id: '%s' Certificate: '%s' Private Key: '%s' CA Certificate: '%s']", commonName, pemFileName, keyFileName, caFileName))

	return nil
}

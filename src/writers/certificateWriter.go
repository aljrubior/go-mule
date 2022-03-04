package writers

import (
	"fmt"
	"github.com/aljrubior/standalone-runtime/security"
	"os"
)

func NewCertificateWriter(privateKey, certificate []byte) CertificateWriter {
	return CertificateWriter{
		privateKey:  privateKey,
		certificate: certificate,
	}
}

type CertificateWriter struct {
	privateKey,
	certificate []byte
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

	pemFile, err := os.Create(pemFileName)

	if err != nil {
		return err
	}

	keyFile, err := os.Create(keyFileName)

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

	println(fmt.Sprintf("Standalone created successfully [Id: '%s' Certificate: '%s' Private Key: '%s']", commonName, pemFileName, keyFileName))

	return nil
}

package persistenceManager

import (
	"fmt"
	"github.com/aljrubior/standalone-runtime/security"
	"os"
)

func NewCertificatePersistanceManager(
	privateKey,
	certificate []byte) CertificatePersistanceManager {
	return CertificatePersistanceManager{}
}

type CertificatePersistanceManager struct {
	privateKey  []byte
	certificate []byte
}

func (persistanceManager CertificatePersistanceManager) WriteFile() error {

	commonName, err := security.NewCertificateWrapper(persistanceManager.certificate).GetCommonName()

	if err != nil {
		return nil
	}

	if err := os.Mkdir(commonName, 644); err != nil {
		return err
	}

	pemFileName := fmt.Sprintf("%s.pem", commonName)
	keyFileName := fmt.Sprintf("%s.key", commonName)

	pemFile, err := os.Create(pemFileName)

	if err != nil {
		return err
	}

	keyFile, err := os.Create(keyFileName)

	if err != nil {
		return err
	}

	_, err = pemFile.Write(persistanceManager.certificate)

	if err != nil {
		return err
	}

	_, err = keyFile.Write(persistanceManager.privateKey)

	return err

}

package security

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func NewTLSConfigBuilder(publicCertificatePath, privateKeyPath, caCertificatePath string) *TLSConfigBuilder {
	return &TLSConfigBuilder{
		PublicCertificatePath: publicCertificatePath,
		PrivateKeyPath:        privateKeyPath,
		CACertificatePath:     caCertificatePath,
	}
}

type TLSConfigBuilder struct {
	PublicCertificatePath string
	PrivateKeyPath        string
	CACertificatePath     string
}

func (builder TLSConfigBuilder) Build() *tls.Config {

	var cert tls.Certificate

	cert.Certificate = append(cert.Certificate, builder.LoadPublicCertificate().Bytes)
	cert.PrivateKey = builder.LoadPrivateKey()

	return &tls.Config{
		InsecureSkipVerify: true,
		MinVersion:         tls.VersionTLS12,
		ClientAuth:         tls.RequireAndVerifyClientCert,
		Certificates:       []tls.Certificate{cert},
		CipherSuites: []uint16{
			tls.TLS_RSA_WITH_AES_128_GCM_SHA256,
			tls.TLS_RSA_WITH_AES_256_CBC_SHA,
			tls.TLS_RSA_WITH_AES_128_CBC_SHA,
		},
	}

}

func (builder TLSConfigBuilder) LoadPublicCertificate() *pem.Block {
	data, err := ioutil.ReadFile(builder.PublicCertificatePath)
	if err != nil {
		println(fmt.Sprintf("ReadFile: %s", builder.PublicCertificatePath))
		log.Println(err)
	}

	block, _ := pem.Decode(data)
	return block
}

func (builder TLSConfigBuilder) LoadPrivateKey() interface{} {

	data, err := ioutil.ReadFile(builder.PrivateKeyPath)

	if err != nil {
		println(fmt.Sprintf("ReadFile: %s", builder.PrivateKeyPath))
		log.Println(err)
	}

	block, _ := pem.Decode(data)

	key, err := x509.ParsePKCS1PrivateKey(block.Bytes)

	if err != nil {
		fmt.Println("Error: ParsePKCS8PrivateKey")
		println(err.Error())
		os.Exit(1)
	}

	return key
}

func (builder TLSConfigBuilder) LoadCACertificate() []byte {
	data, err := ioutil.ReadFile(builder.CACertificatePath)
	if err != nil {
		println(fmt.Sprintf("ReadFile: %s", builder.PublicCertificatePath))
		log.Println(err)
	}

	return data
}

func (builder TLSConfigBuilder) BuildCertPool(caCertificate []byte) *x509.CertPool {

	pool := x509.NewCertPool()

	ok := pool.AppendCertsFromPEM(caCertificate)

	if !ok {
		println("Error: AppendCertsFromPEM")
	}

	return pool
}

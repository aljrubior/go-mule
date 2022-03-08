package security

import "github.com/aljrubior/go-mule/errors"

func (generator RSAKeyPairGenerator) ThrowRSAKeyPairGeneratorError(message string, err error) *errors.RSAKeyPairGeneratorError {

	return errors.NewRSAKeyPairGeneratorError(message, err.Error())
}

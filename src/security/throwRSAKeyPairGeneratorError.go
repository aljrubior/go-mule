package security

import "github.com/aljrubior/standalone-runtime/errors"

func (generator RSAKeyPairGenerator) ThrowRSAKeyPairGeneratorError(message string, err error) *errors.RSAKeyPairGeneratorError {

	return errors.NewRSAKeyPairGeneratorError(message, err.Error())
}

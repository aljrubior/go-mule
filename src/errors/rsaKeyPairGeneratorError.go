package errors

func NewRSAKeyPairGeneratorError(message, reason string) *RSAKeyPairGeneratorError {
	return &RSAKeyPairGeneratorError{
		message: message,
		reason:  reason,
	}
}

type RSAKeyPairGeneratorError struct {
	message, reason string
}

func (this *RSAKeyPairGeneratorError) Error() string {
	return this.message
}

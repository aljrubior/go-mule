package serverRegistrationManager

import "strings"

func NewCSRWrapper(csr []byte) CSRWrapper {
	return CSRWrapper{
		csr: csr,
	}
}

type CSRWrapper struct {
	csr []byte
}

func (wrapper CSRWrapper) ToString() string {

	csrAsString := string(wrapper.csr)

	csrAsArray := strings.Split(csrAsString, "\n")

	return strings.Join(csrAsArray, "\n")
}

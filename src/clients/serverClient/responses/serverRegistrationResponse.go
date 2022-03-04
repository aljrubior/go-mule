package responses

type ServerRegistrationResponse struct {
	Certificate     string          `json:"certificate"`
	CACertificate   string          `json:"caCertificate"`
	EnvironmentUrls EnvironmentUrls `json:"environmentUrls"`
}

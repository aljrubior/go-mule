package responses

type ServerRegistrationResponse struct {
	Certificate     string          `json:"certificate"`
	EnvironmentUrls EnvironmentUrls `json:"environmentUrls"`
}

package entities

func NewServerEntity(
	privateKey,
	certificate,
	caCertificate []byte,
	websocketUrl,
	metricsUrl string) *ServerEntity {

	return &ServerEntity{
		PrivateKey:    privateKey,
		Certificate:   certificate,
		CACertificate: caCertificate,
		WebsocketUrl:  websocketUrl,
		MetricsUrl:    metricsUrl,
	}
}

type ServerEntity struct {
	PrivateKey    []byte
	Certificate   []byte
	CACertificate []byte
	WebsocketUrl  string
	MetricsUrl    string
}

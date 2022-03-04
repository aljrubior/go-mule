package entities

func NewServerEntity(
	privateKey,
	certificate []byte,
	websocketUrl,
	metricsUrl string) *ServerEntity {

	return &ServerEntity{
		PrivateKey:   privateKey,
		Certificate:  certificate,
		WebsocketUrl: websocketUrl,
		MetricsUrl:   metricsUrl,
	}
}

type ServerEntity struct {
	PrivateKey   []byte
	Certificate  []byte
	WebsocketUrl string
	MetricsUrl   string
}

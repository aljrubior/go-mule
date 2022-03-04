package notifications

func NewPostHandShakeNotification() *PostHandShakeNotification {
	return &PostHandShakeNotification{}
}

type PostHandShakeNotification struct {
	message string
}

func (this PostHandShakeNotification) CreateNotification() string {
	return `POST handshake HTTP/1.1
Content-Type: application/json
Message-Id: 382b99bf-640c-4e7e-b1f4-00a8e9da1b3c
accept: application/json
Content-length:689

{"addresses":"[{\"ip\":\"192.168.1.15\",\"networkInterface\":\"eth1\"},{\"ip\":\"10.0.2.15\",\"networkInterface\":\"eth0\"},{\"ip\":\"127.0.0.1\",\"networkInterface\":\"lo\"}]","muleVersion":"4.3.0","serverType":"GATEWAY","agentVersion":"2.4.28-SNAPSHOT","timeZone":"UTC","gatewayVersion":"4.3.0","runtimeInformation":"{\"osInformation\":{\"name\":\"Linux\",\"version\":\"4.15.0-58-generic\",\"architecture\":\"amd64\"},\"jvmInformation\":{\"runtime\":{\"name\":\"OpenJDK Runtime Environment\",\"version\":\"1.8.0_292-b10\"},\"specification\":{\"vendor\":\"AdoptOpenJDK\",\"name\":\"Java Platform API Specification\",\"version\":\"1.8\"}},\"muleLicenseExpirationDate\":\"1648512000000\"}"}`
}

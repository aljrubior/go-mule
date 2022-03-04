package runtime

import (
	"github.com/aljrubior/standalone-runtime/runtime/messages"
	"github.com/aljrubior/standalone-runtime/runtime/messages/notifications"
	"github.com/aljrubior/standalone-runtime/runtime/messages/responses"
	"github.com/aljrubior/standalone-runtime/tls"
	"github.com/aljrubior/standalone-runtime/websockets"
	"github.com/gorilla/websocket"
	"io"
	"log"
	"net/url"
	"time"
)

func NewStandaloneRuntime(
	serverId,
	contextId,
	certificatePath,
	privateKeyPath,
	caCertificate string) StandaloneRuntime {

	return StandaloneRuntime{
		serverId:        serverId,
		contextId:       contextId,
		certificatePath: certificatePath,
		privateKeyPath:  privateKeyPath,
		caCertificate:   caCertificate,
	}
}

type StandaloneRuntime struct {
	serverId        string
	contextId       string
	certificatePath string
	privateKeyPath  string
	caCertificate   string
}

func (runtime StandaloneRuntime) Start() {

	conn := runtime.createWebsockerConnection()

	defer conn.Close()

	go runtime.startMessageListener(conn)

	runtime.initiateAnypointHandShake(conn)

	runtime.sendStartupNotifications(conn)

	runtime.startKeepAliveNotifications(conn)
}

func (runtime StandaloneRuntime) initiateAnypointHandShake(conn *websocket.Conn) {
	notification := notifications.NewPostHandShakeNotification().CreateNotification()
	runtime.SendNotification(conn, notification)
}

func (runtime StandaloneRuntime) sendStartupNotifications(conn *websocket.Conn) {
	notification := notifications.NewPostHandShakeNotification().CreateNotification()
	runtime.SendNotification(conn, notification)
}

func (runtime StandaloneRuntime) startKeepAliveNotifications(conn *websocket.Conn) {
	notification := notifications.NewPostKeepAliveNotification().CreateNotification()

	for {
		runtime.SendNotification(conn, notification)

		time.Sleep(2 * time.Second)
	}
}

func (runtime StandaloneRuntime) SendNotification(conn *websocket.Conn, notification string) {
	if err := conn.WriteMessage(websocket.BinaryMessage, []byte(notification)); err != nil {
		log.Fatal(err)
	}
}

func (runtime StandaloneRuntime) createWebsockerConnection() *websocket.Conn {

	url := runtime.CreateURL()

	tlsConfig := tls.NewTLSConfigBuilder(runtime.certificatePath, runtime.privateKeyPath, runtime.caCertificate).Build()

	conn, _, err := websockets.NewRuntimeManagerDialer(url, tlsConfig).CreateDialer()

	if err != nil {
		println("Dial...")
		log.Fatalln(err)
	}

	return conn
}

func (runtime StandaloneRuntime) CreateURL() url.URL {
	return url.URL{
		Scheme: "wss",
		Host:   "runtime-manager.qax.anypoint.mulesoft.com:443",
		Path:   "/mule",
	}
}

func (runtime StandaloneRuntime) startMessageListener(conn *websocket.Conn) {

	for {
		_, message, err := conn.ReadMessage()

		if err != nil || err == io.EOF {
			log.Fatal("Error reading: ", err)
			break
		}

		websocketMessage := messages.NewWebsocketMessage(string(message))
		println(websocketMessage.GetMessage())

		switch websocketMessage.GetResquestHeader() {
		case messages.GET_CLUSTERS_REQUEST_HEADER:
			message := responses.NewGetClustersResponse(websocketMessage.GetMessageId()).CreateResponse()
			println(message)
			if err := conn.WriteMessage(websocket.BinaryMessage, []byte(message)); err != nil {
				println("createGetClustersResponse")
				log.Fatal(err)
			}
		case messages.GET_AGENT_CONFIGURATION_REQUEST_HEADER:
			responses.NewAgentConfigurationResponse(websocketMessage.GetMessageId()).CreateResponse()
			if err := conn.WriteMessage(websocket.BinaryMessage, []byte(message)); err != nil {
				println("createAgentConfigurationResponse")
				log.Fatal(err)
			}
		}

		notification := notifications.NewPutDomainDeploymentContextInitialisedNotification(runtime.serverId, runtime.contextId).CreateNotification()

		if err := conn.WriteMessage(websocket.BinaryMessage, []byte(notification)); err != nil {
			println("createPutDomainDeploymentContextInitialisedMessage")
			log.Fatal(err)
		}

		time.Sleep(1 * time.Second)

		notification = notifications.NewPutDomainDeploymentDeployedNotification(runtime.serverId, runtime.contextId).CreateNotification()
		if err := conn.WriteMessage(websocket.BinaryMessage, []byte(notification)); err != nil {
			println("createPutDomainDeploymentDeployedMessage")
			log.Fatal(err)
		}

		time.Sleep(1 * time.Second)

		notification = notifications.NewPutStartupNotification(runtime.serverId, runtime.contextId).CreateNotification()

		if err := conn.WriteMessage(websocket.BinaryMessage, []byte(notification)); err != nil {
			println("createPutStartupMessage")
			log.Fatal(err)
		}
	}
}

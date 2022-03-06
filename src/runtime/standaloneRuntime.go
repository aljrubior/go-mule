package runtime

import (
	"fmt"
	"github.com/aljrubior/standalone-runtime/application"
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
	caCertificatePath string) StandaloneRuntime {

	return StandaloneRuntime{
		serverId:          serverId,
		contextId:         contextId,
		certificatePath:   certificatePath,
		privateKeyPath:    privateKeyPath,
		caCertificatePath: caCertificatePath,
	}
}

type StandaloneRuntime struct {
	serverId          string
	contextId         string
	certificatePath   string
	privateKeyPath    string
	caCertificatePath string
}

func (runtime StandaloneRuntime) Start() {

	conn := runtime.createWebsocketConnection()

	defer conn.Close()

	go runtime.startMessageListener(conn)

	runtime.initiateAnypointHandShake(conn)

	runtime.sendDomainContextInitialisedNotifications(conn)

	runtime.sendDomainDeployedNotifications(conn)

	runtime.sendStartupNotifications(conn)

	runtime.startKeepAliveNotifications(conn)
}

func (runtime StandaloneRuntime) initiateAnypointHandShake(conn *websocket.Conn) {
	notification := notifications.NewPostHandShakeNotification().CreateNotification()
	runtime.SendNotification(conn, notification)
}

func (runtime StandaloneRuntime) sendDomainContextInitialisedNotifications(conn *websocket.Conn) {
	notification := notifications.NewPutDomainDeploymentContextInitialisedNotification(runtime.serverId, runtime.contextId).CreateNotification()
	runtime.SendNotification(conn, notification)
}

func (runtime StandaloneRuntime) sendDomainDeployedNotifications(conn *websocket.Conn) {
	notification := notifications.NewPutDomainDeploymentDeployedNotification(runtime.serverId, runtime.contextId).CreateNotification()
	runtime.SendNotification(conn, notification)
}

func (runtime StandaloneRuntime) sendStartupNotifications(conn *websocket.Conn) {
	notification := notifications.NewPutStartupNotification(runtime.serverId, runtime.contextId).CreateNotification()
	runtime.SendNotification(conn, notification)
}

func (runtime StandaloneRuntime) startKeepAliveNotifications(conn *websocket.Conn) {
	notification := notifications.NewPostKeepAliveNotification().CreateNotification()

	for {
		runtime.SendNotification(conn, notification)

		time.Sleep(1 * time.Second)
	}
}

func (runtime StandaloneRuntime) SendNotification(conn *websocket.Conn, notification string) {

	println(notification)

	if err := conn.WriteMessage(websocket.BinaryMessage, []byte(notification)); err != nil {
		log.Fatal(err)
	}

	time.Sleep(1 * time.Second)
}

func (runtime StandaloneRuntime) createWebsocketConnection() *websocket.Conn {

	url := runtime.CreateURL()

	tlsConfig := tls.NewTLSConfigBuilder(runtime.certificatePath, runtime.privateKeyPath, runtime.caCertificatePath).Build()

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

		if websocketMessage.IsResponse() {
			println(fmt.Sprintf("\n%s", websocketMessage.GetMessage()))
			continue
		}

		println(websocketMessage.GetResquestHeader())

		switch websocketMessage.GetRequestAction() {
		case messages.GET_CLUSTERS_REQUEST_ACTION:

			println(fmt.Sprintf("\n%s", websocketMessage.GetMessage()))

			message := responses.NewGetClustersResponse(websocketMessage.GetMessageId()).CreateResponse()

			println(message)

			if err := conn.WriteMessage(websocket.BinaryMessage, []byte(message)); err != nil {
				log.Fatal(err)
			}
		case messages.GET_AGENT_CONFIGURATION_REQUEST_ACTION:

			println(fmt.Sprintf("\n%s", websocketMessage.GetMessage()))

			message := responses.NewAgentConfigurationResponse(websocketMessage.GetMessageId()).CreateResponse()

			println(message)

			if err := conn.WriteMessage(websocket.BinaryMessage, []byte(message)); err != nil {
				log.Fatal(err)
			}
		case messages.PUT_APPLICATIONS_REQUEST_ACTION:
			applicationName := websocketMessage.GetApplicationName()

			totalFixedScheduler := 50
			totalCronScheduler := 50

			application := application.NewApplicationBuilder(applicationName, totalFixedScheduler, totalCronScheduler).Build()

			message := notifications.NewPutDeploymentStartedNotification(runtime.serverId, runtime.contextId, application).CreateNotification()
			runtime.SendNotification(conn, message)

			message = notifications.NewPutDeploymentContextCreatedNotification(runtime.serverId, runtime.contextId, application).CreateNotification()
			runtime.SendNotification(conn, message)

			message = notifications.NewPutDeploymentContextInitialisedNotification(runtime.serverId, runtime.contextId, application).CreateNotification()
			runtime.SendNotification(conn, message)

			time.Sleep(1 * time.Second)

			messages := notifications.NewPutDeploymentFlowNotification(runtime.serverId, runtime.contextId, application).CreateNotifications()
			for _, v := range messages {
				runtime.SendNotification(conn, v)
			}

			message = notifications.NewPutDeploymentContextStartedNotification(runtime.serverId, runtime.contextId, application).CreateNotification()
			runtime.SendNotification(conn, message)

			message = notifications.NewPutDeploymentSchedulersNotification(runtime.serverId, runtime.contextId, application).CreateNotification()
			runtime.SendNotification(conn, message)

			message = notifications.NewPutDeploymentDeployedNotification(runtime.serverId, runtime.contextId, application).CreateNotification()
			runtime.SendNotification(conn, message)
		}
	}
}

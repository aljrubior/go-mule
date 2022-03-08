package runtime

import (
	"crypto/tls"
	"fmt"
	"github.com/aljrubior/standalone-runtime/application"
	"github.com/aljrubior/standalone-runtime/managers/metricManager"
	"github.com/aljrubior/standalone-runtime/runtime/messages"
	"github.com/aljrubior/standalone-runtime/runtime/messages/notifications"
	"github.com/aljrubior/standalone-runtime/runtime/strategies"
	"github.com/aljrubior/standalone-runtime/websockets"
	"github.com/gorilla/websocket"
	"io"
	"log"
	"net/url"
	"time"
)

func NewStandaloneRuntime(
	serverId,
	contextId string,
	tlsConfig *tls.Config,
	totalFixedSchPerApp,
	totalCronSchPerApp int) StandaloneRuntime {

	applications := make(map[string]*application.Application)

	return StandaloneRuntime{
		serverId:            serverId,
		contextId:           contextId,
		tlsConfig:           tlsConfig,
		totalFixedSchPerApp: totalFixedSchPerApp,
		totalCronSchPerApp:  totalCronSchPerApp,

		applications: applications,
	}
}

type StandaloneRuntime struct {
	serverId,
	contextId string
	tlsConfig *tls.Config
	totalFixedSchPerApp,
	totalCronSchPerApp int
	metricManager metricManager.MetricManager

	applications map[string]*application.Application
}

func (t StandaloneRuntime) GetApplications() *map[string]*application.Application {
	return &t.applications
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

	conn, _, err := websockets.NewRuntimeManagerDialer(url, runtime.tlsConfig).CreateDialer()

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

func (t StandaloneRuntime) startMessageListener(conn *websocket.Conn) {

	regex := messages.NewActionRequestRegex()

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

		if regex.PutApplication.MatchString(websocketMessage.GetResquestHeader()) {
			applicationName := websocketMessage.GetApplicationName()
			t.applications[applicationName] = application.NewApplicationBuilder(applicationName, t.totalFixedSchPerApp, t.totalCronSchPerApp).Build()
		}

		strategies.NewActionStrategyBuilder(conn, websocketMessage, t.serverId, t.contextId, &t.applications, &regex).Build().Execute()

	}
}

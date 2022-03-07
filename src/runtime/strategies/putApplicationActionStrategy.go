package strategies

import (
	"github.com/aljrubior/standalone-runtime/application"
	"github.com/aljrubior/standalone-runtime/runtime/messages"
	"github.com/aljrubior/standalone-runtime/runtime/messages/notifications"
	"github.com/gorilla/websocket"
	"time"
)

func NewPutApplicationActionStrategy(
	conn *websocket.Conn,
	message *messages.WebsocketMessage,
	serverId,
	contextId string,
	application *application.Application) *PutApplicationActionStrategy {
	return &PutApplicationActionStrategy{
		conn:        conn,
		message:     message,
		serverId:    serverId,
		contextId:   contextId,
		application: application,
	}
}

type PutApplicationActionStrategy struct {
	BaseActionStrategy
	conn    *websocket.Conn
	message *messages.WebsocketMessage
	serverId,
	contextId string
	application *application.Application
}

func (t *PutApplicationActionStrategy) Execute() {

	message := notifications.NewPutDeploymentStartedNotification(t.serverId, t.contextId, t.application).CreateNotification()
	t.SendMessage(t.conn, message)

	message = notifications.NewPutDeploymentContextCreatedNotification(t.serverId, t.contextId, t.application).CreateNotification()
	t.SendMessage(t.conn, message)

	message = notifications.NewPutDeploymentContextInitialisedNotification(t.serverId, t.contextId, t.application).CreateNotification()
	t.SendMessage(t.conn, message)

	time.Sleep(1 * time.Second)

	messages := notifications.NewPutDeploymentFlowNotification(t.serverId, t.contextId, t.application, "STARTED").CreateNotifications()

	for _, v := range messages {
		t.SendMessage(t.conn, v)
	}

	message = notifications.NewPutDeploymentContextStartedNotification(t.serverId, t.contextId, t.application).CreateNotification()
	t.SendMessage(t.conn, message)

	message = notifications.NewPutDeploymentSchedulersNotification(t.serverId, t.contextId, t.application).CreateNotification()
	t.SendMessage(t.conn, message)

	message = notifications.NewPutDeploymentDeployedNotification(t.serverId, t.contextId, t.application).CreateNotification()
	t.SendMessage(t.conn, message)
}

package strategies

import (
	"github.com/aljrubior/go-mule/application"
	"github.com/aljrubior/go-mule/runtime/messages"
	"github.com/aljrubior/go-mule/runtime/messages/notifications"
	"github.com/gorilla/websocket"
)

func NewPutStartApplicationActionStrategy(
	conn *websocket.Conn,
	message *messages.WebsocketMessage,
	serverId,
	contextId string,
	application *application.Application) *PutStartApplicationActionStrategy {
	return &PutStartApplicationActionStrategy{
		conn:        conn,
		message:     message,
		serverId:    serverId,
		contextId:   contextId,
		application: application,
	}
}

type PutStartApplicationActionStrategy struct {
	BaseActionStrategy
	conn    *websocket.Conn
	message *messages.WebsocketMessage
	serverId,
	contextId string
	application *application.Application
}

func (t *PutStartApplicationActionStrategy) Execute() {

	messages := notifications.NewPutDeploymentFlowNotification(t.serverId, t.contextId, t.application, "STARTED").CreateNotifications()

	for _, v := range messages {
		t.SendMessage(t.conn, v)
	}

	message := notifications.NewPutDeploymentContextStartedNotification(t.serverId, t.contextId, t.application).CreateNotification()
	t.SendMessage(t.conn, message)
}

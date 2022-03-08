package strategies

import (
	"github.com/aljrubior/go-mule/application"
	"github.com/aljrubior/go-mule/runtime/messages"
	"github.com/aljrubior/go-mule/runtime/messages/notifications"
	"github.com/gorilla/websocket"
)

func NewPutStopApplicationActionStrategy(
	conn *websocket.Conn,
	message *messages.WebsocketMessage,
	serverId,
	contextId string,
	application *application.Application) *PutStopApplicationActionStrategy {
	return &PutStopApplicationActionStrategy{
		conn:        conn,
		message:     message,
		serverId:    serverId,
		contextId:   contextId,
		application: application,
	}
}

type PutStopApplicationActionStrategy struct {
	BaseActionStrategy
	conn    *websocket.Conn
	message *messages.WebsocketMessage
	serverId,
	contextId string
	application *application.Application
}

func (t *PutStopApplicationActionStrategy) Execute() {

	messages := notifications.NewPutDeploymentFlowNotification(t.serverId, t.contextId, t.application, "STOPPED").CreateNotifications()

	for _, v := range messages {
		t.SendMessage(t.conn, v)
	}

	message := notifications.NewPutDeploymentContextStoppedNotification(t.serverId, t.contextId, t.application).CreateNotification()
	t.SendMessage(t.conn, message)
}

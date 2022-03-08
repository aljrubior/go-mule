package strategies

import (
	"fmt"
	"github.com/aljrubior/go-mule/runtime/messages"
	"github.com/aljrubior/go-mule/runtime/messages/responses"
	"github.com/gorilla/websocket"
)

func NewGetClusterActionStrategy(
	conn *websocket.Conn,
	message *messages.WebsocketMessage) *GetClusterActionStrategy {
	return &GetClusterActionStrategy{
		conn:    conn,
		message: message,
	}
}

type GetClusterActionStrategy struct {
	BaseActionStrategy
	conn    *websocket.Conn
	message *messages.WebsocketMessage
}

func (t *GetClusterActionStrategy) Execute() {

	println(fmt.Sprintf("\n%s", t.message.GetMessage()))

	response := responses.NewGetClustersResponse(t.message.GetMessageId()).CreateResponse()

	t.SendMessage(t.conn, response)
}

package strategies

import (
	"fmt"
	"github.com/aljrubior/go-mule/runtime/messages"
	"github.com/aljrubior/go-mule/runtime/messages/responses"
	"github.com/gorilla/websocket"
)

func NewGetAgentConfigurationActionStrategy(
	conn *websocket.Conn,
	message *messages.WebsocketMessage) *GetAgentConfigurationActionStrategy {
	return &GetAgentConfigurationActionStrategy{
		conn:    conn,
		message: message,
	}
}

type GetAgentConfigurationActionStrategy struct {
	BaseActionStrategy
	conn    *websocket.Conn
	message *messages.WebsocketMessage
}

func (t *GetAgentConfigurationActionStrategy) Execute() {

	println(fmt.Sprintf("\n%s", t.message.GetMessage()))

	response := responses.NewGetClustersResponse(t.message.GetMessageId()).CreateResponse()

	t.SendMessage(t.conn, response)
}

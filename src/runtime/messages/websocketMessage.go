package messages

import "strings"

const (
	MESSAGE_ID_HEADER                          = "Message-Id"
	GET_CLUSTERS_REQUEST_HEADER                = "GET clusters HTTP/1.1"
	GET_AGENT_CONFIGURATION_REQUEST_HEADER     = "GET agent/configuration HTTP/1.1"
	PUT_APPLICATIONS_REQUEST_HEADER            = "PUT applications/app-noports-mule4-only HTTP/1.1"
	PATCH_LOGGING_SERVICE_REQUEST_HEADER       = "PATCH agent/mule.agent.logging.service HTTP/1.1"
	PUT_LOGGING_SERVICE_ENABLE_REQUEST_HEADER  = "PUT agent/mule.agent.logging.service/enable HTTP/1.1"
	PUT_TRACKING_SERVICE_REQUEST_HEADER        = "PUT agent/mule.agent.tracking.service HTTP/1.1"
	PUT_TRACKING_SERVICE_ENABLE_REQUEST_HEADER = "PUT agent/mule.agent.tracking.service/enable HTTP/1.1"
)

func NewWebsocketMessage(message string) *WebsocketMessage {

	return &WebsocketMessage{
		message: message,
	}
}

type WebsocketMessage struct {
	message   string
	messageId string
}

func (this WebsocketMessage) GetResquestHeader() string {

	return this.message[:strings.IndexByte(this.message, '\n')-1]
}

func (this WebsocketMessage) GetMessageId() string {

	if this.messageId == "" {
		index := strings.Index(this.message, MESSAGE_ID_HEADER)

		if index < 0 {
			return ""
		}

		messageId := this.message[index:]
		messageId = messageId[:strings.IndexByte(messageId, '\n')-1]
		result := strings.Split(messageId, ":")
		return strings.TrimSpace(result[1])

	}

	return this.messageId
}

func (this WebsocketMessage) GetMessage() string {
	return this.message
}

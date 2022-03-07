package messages

import (
	"strings"
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

func (this WebsocketMessage) IsResponse() bool {

	return strings.HasPrefix(this.message, "HTTP/1.1")
}

func (this WebsocketMessage) IsRequest() bool {

	return !strings.HasPrefix(this.message, "HTTP/1.1")
}

func (this WebsocketMessage) IsDeployApplicationRequest() bool {

	return strings.HasPrefix(this.GetResquestHeader(), "PUT applications/")
}

func (this WebsocketMessage) GetApplicationName() string {

	if this.IsDeployApplicationRequest() {
		request := strings.Split(this.GetResquestHeader(), " ")
		return strings.Split(request[1], "/")[1]
	}
	return ""
}

func (this WebsocketMessage) GetRequestAction() string {

	header := this.GetResquestHeader()

	if this.IsResponse() {
		return header
	}

	if this.IsDeployApplicationRequest() {
		return header[:strings.IndexByte(header, '/')+1]
	}

	return header[:strings.Index(header, " HTTP/1.1")]
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

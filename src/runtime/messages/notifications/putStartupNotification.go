package notifications

import "strings"

func NewPutStartupNotification(serverId, contextId string) *PutStartupNotification {
	return &PutStartupNotification{
		Notification: Notification{
			ServerId:  serverId,
			ContextId: contextId,
		},
	}
}

type PutStartupNotification struct {
	Notification
}

func (notification PutStartupNotification) CreateNotification() string {
	message := notification.GetTemplate()

	message = strings.Replace(message, SERVER_ID_PATTERN, notification.ServerId, 1)
	message = strings.Replace(message, CONTEXT_ID_PATTERN, notification.ContextId, 1)

	return message
}

func (notification PutStartupNotification) GetTemplate() string {
	return `POST startup HTTP/1.1
Content-Type: application/json
X-ANYPNT-SERVER-ID: {{serverId}}
X-ANYPNT-CTX-ID: {{contextId}}
X-ANYPNT-AGENT-VERSION: 2.4.28-SNAPSHOT
Content-Length: 234
Content-Encoding: UTF-8

{"addresses":"[{\"ip\":\"192.168.1.15\",\"networkInterface\":\"eth1\"},{\"ip\":\"10.0.2.15\",\"networkInterface\":\"eth0\"},{\"ip\":\"127.0.0.1\",\"networkInterface\":\"lo\"}]","domains":"[{\"applications\":[],\"name\":\"default\"}]"}`
}

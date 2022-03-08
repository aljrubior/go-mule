package notifications

import (
	"github.com/aljrubior/go-mule/application"
	"strings"
)

func NewPutDeploymentContextStoppedNotification(serverId, contextId string, application *application.Application) *PutDeploymentContextStoppedNotification {
	return &PutDeploymentContextStoppedNotification{
		ApplicationNotification: ApplicationNotification{
			Notification: Notification{
				ServerId:  serverId,
				ContextId: contextId,
			},
			Application: application,
		},
	}
}

type PutDeploymentContextStoppedNotification struct {
	ApplicationNotification
}

func (notification PutDeploymentContextStoppedNotification) CreateNotification() string {
	message := notification.GetTemplate()

	message = strings.ReplaceAll(message, APPLICATION_NAME_PATTERN, notification.Application.Name)
	message = strings.ReplaceAll(message, SERVER_ID_PATTERN, notification.ServerId)
	message = strings.ReplaceAll(message, CONTEXT_ID_PATTERN, notification.ContextId)

	return message
}

func (notification PutDeploymentContextStoppedNotification) GetTemplate() string {
	return `PUT applications/{{applicationName}}/deployment HTTP/1.1
Content-Type: application/json
X-ANYPNT-SERVER-ID: {{serverId}}
X-ANYPNT-CTX-ID: {{contextId}}
X-ANYPNT-AGENT-VERSION: 2.4.28-SNAPSHOT
Content-Length: 1000
Content-Encoding: UTF-8

{"status":"CONTEXT_STOPPED","message":"","application":{"name":"{{applicationName}}","domain":"default","state":"STOPPED","flows":null,"lastDateStarted":null}}`
}


package notifications

import (
	"github.com/aljrubior/standalone-runtime/application"
	"strings"
)

func NewPutDeploymentStartedNotification(serverId, contextId string, application *application.Application) *PutDeploymentStartedNotification {
	return &PutDeploymentStartedNotification{
		ApplicationNotification: ApplicationNotification{
			Notification: Notification{
				ServerId:  serverId,
				ContextId: contextId,
			},
			Application: application,
		},
	}
}

type PutDeploymentStartedNotification struct {
	ApplicationNotification
}

func (notification PutDeploymentStartedNotification) CreateNotification() string {
	message := notification.GetTemplate()

	message = strings.ReplaceAll(message, SERVER_ID_PATTERN, notification.ServerId)
	message = strings.ReplaceAll(message, CONTEXT_ID_PATTERN, notification.ContextId)
	message = strings.ReplaceAll(message, APPLICATION_NAME_PATTERN, notification.Application.Name)

	return message
}

func (notification PutDeploymentStartedNotification) GetTemplate() string {
	return `PUT applications/{{applicationName}}/deployment HTTP/1.1
Content-Type: application/json
X-ANYPNT-SERVER-ID: {{serverId}}
X-ANYPNT-CTX-ID: {{contextId}}
X-ANYPNT-AGENT-VERSION: 2.4.28-SNAPSHOT
Content-Length: 1000
Content-Encoding: UTF-8

{"status":"DEPLOYMENT_STARTED","message":"","application":{"name":"{{applicationName}}","domain":"default","state":"CREATED","flows":null,"lastDateStarted":null}}`
}

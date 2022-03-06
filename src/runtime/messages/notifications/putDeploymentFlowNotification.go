package notifications

import (
	"github.com/aljrubior/standalone-runtime/application"
	"strings"
)

const (
	FLOW_NAME_PATTERN = "{{flowName}}"
)

func NewPutDeploymentFlowNotification(serverId, contextId string, application *application.Application) *PutDeploymentFlowNotification {
	return &PutDeploymentFlowNotification{
		ApplicationNotification: ApplicationNotification{
			Notification: Notification{
				ServerId:  serverId,
				ContextId: contextId,
			},
			Application: application,
		},
	}
}

type PutDeploymentFlowNotification struct {
	ApplicationNotification
}

func (notification PutDeploymentFlowNotification) CreateNotifications() []string {
	var notifications []string

	for _, v := range notification.Application.FixedSchedulers {
		notifications = append(notifications, notification.createNotification(notification.ServerId, notification.ContextId, notification.Application.Name, v.FlowName))
	}

	for _, v := range notification.Application.CronSchedulers {
		notifications = append(notifications, notification.createNotification(notification.ServerId, notification.ContextId, notification.Application.Name, v.FlowName))
	}

	return notifications
}

func (notification PutDeploymentFlowNotification) createNotification(serverId, contextId, applicationName, flowName string) string {
	message := notification.GetTemplate()

	message = strings.ReplaceAll(message, SERVER_ID_PATTERN, serverId)
	message = strings.ReplaceAll(message, CONTEXT_ID_PATTERN, contextId)
	message = strings.ReplaceAll(message, APPLICATION_NAME_PATTERN, applicationName)
	message = strings.ReplaceAll(message, FLOW_NAME_PATTERN, flowName)

	return message
}

func (notification PutDeploymentFlowNotification) GetTemplate() string {
	return `PUT applications/{{applicationName}}/flows/{{flowName}} HTTP/1.1
Content-Type: application/json
X-ANYPNT-SERVER-ID: {{serverId}}
X-ANYPNT-CTX-ID: {{contextId}}
X-ANYPNT-AGENT-VERSION: 2.4.28-SNAPSHOT
Content-Length: 1000
Content-Encoding: UTF-8

{"applicationName":"{{applicationName}}","flow":{"name":"{{flowName}}","status":"STARTED","defaultStatus":"STARTED"}}`
}

package notifications

import (
	"github.com/aljrubior/standalone-runtime/application"
	"strings"
)

const (
	FLOW_NAME_PATTERN   = "{{flowName}}"
	FLOW_STATUS_PATTERN = "{{flowStatus}}"
)

func NewPutDeploymentFlowNotification(
	serverId,
	contextId string,
	application *application.Application,
	flowStatus string) *PutDeploymentFlowNotification {
	return &PutDeploymentFlowNotification{
		ApplicationNotification: ApplicationNotification{
			Notification: Notification{
				ServerId:  serverId,
				ContextId: contextId,
			},
			Application: application,
		},
		FlowStatus: flowStatus,
	}
}

type PutDeploymentFlowNotification struct {
	ApplicationNotification
	FlowStatus string
}

func (t PutDeploymentFlowNotification) CreateNotifications() []string {
	var notifications []string

	for _, v := range t.Application.FixedSchedulers {
		notifications = append(notifications, t.createNotification(t.ServerId, t.ContextId, t.Application.Name, v.FlowName))
	}

	for _, v := range t.Application.CronSchedulers {
		notifications = append(notifications, t.createNotification(t.ServerId, t.ContextId, t.Application.Name, v.FlowName))
	}

	return notifications
}

func (t PutDeploymentFlowNotification) createNotification(serverId, contextId, applicationName, flowName string) string {
	message := t.GetTemplate()

	message = strings.ReplaceAll(message, SERVER_ID_PATTERN, serverId)
	message = strings.ReplaceAll(message, CONTEXT_ID_PATTERN, contextId)
	message = strings.ReplaceAll(message, APPLICATION_NAME_PATTERN, applicationName)
	message = strings.ReplaceAll(message, FLOW_NAME_PATTERN, flowName)
	message = strings.ReplaceAll(message, FLOW_STATUS_PATTERN, t.FlowStatus)

	return message
}

func (t PutDeploymentFlowNotification) GetTemplate() string {
	return `PUT applications/{{applicationName}}/flows/{{flowName}} HTTP/1.1
Content-Type: application/json
X-ANYPNT-SERVER-ID: {{serverId}}
X-ANYPNT-CTX-ID: {{contextId}}
X-ANYPNT-AGENT-VERSION: 2.4.28-SNAPSHOT
Content-Length: 1000
Content-Encoding: UTF-8

{"applicationName":"{{applicationName}}","flow":{"name":"{{flowName}}","status":"{{flowStatus}}","defaultStatus":"STARTED"}}`
}

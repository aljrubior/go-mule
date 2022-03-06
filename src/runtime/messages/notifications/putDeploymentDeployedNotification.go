package notifications

import (
	"fmt"
	"github.com/aljrubior/standalone-runtime/application"
	"strings"
)

func NewPutDeploymentDeployedNotification(serverId, contextId string, application *application.Application) *PutDeploymentDeployedNotification {
	return &PutDeploymentDeployedNotification{
		ApplicationNotification: ApplicationNotification{
			Notification: Notification{
				ServerId:  serverId,
				ContextId: contextId,
			},
			Application: application,
		},
	}
}

type PutDeploymentDeployedNotification struct {
	ApplicationNotification
}

func (notification PutDeploymentDeployedNotification) CreateNotification() string {
	message := notification.GetTemplate()

	message = strings.ReplaceAll(message, SERVER_ID_PATTERN, notification.ServerId)
	message = strings.ReplaceAll(message, CONTEXT_ID_PATTERN, notification.ContextId)
	message = strings.ReplaceAll(message, APPLICATION_NAME_PATTERN, notification.Application.Name)

	flows := notification.BuildFlows()
	message = strings.ReplaceAll(message, FLOWS_PATTERN, flows)

	return message
}

func (notification PutDeploymentDeployedNotification) BuildFlows() string {
	template := notification.GetFlowTemplate()
	var sb strings.Builder

	sep := ""

	for _, v := range notification.Application.FixedSchedulers {
		sb.WriteString(fmt.Sprintf("%s%s", sep, strings.ReplaceAll(template, FLOW_NAME_PATTERN, v.FlowName)))
		sep = ","
	}

	for _, v := range notification.Application.CronSchedulers {
		sb.WriteString(fmt.Sprintf("%s%s", sep, strings.ReplaceAll(template, FLOW_NAME_PATTERN, v.FlowName)))
	}

	return sb.String()
}

func (notification PutDeploymentDeployedNotification) GetTemplate() string {
	return `PUT applications/{{applicationName}}/deployment HTTP/1.1
Content-Type: application/json
X-ANYPNT-SERVER-ID: {{serverId}}
X-ANYPNT-CTX-ID: {{contextId}}
X-ANYPNT-AGENT-VERSION: 2.4.28-SNAPSHOT
Content-Length: 40000
Content-Encoding: UTF-8

{"status":"DEPLOYED","message":"","application":{"name":"{{applicationName}}","domain":"default","state":"STARTED","flows":[{{flows}}],"lastDateStarted":1646568892702}}`
}

func (notification PutDeploymentDeployedNotification) GetFlowTemplate() string {
	return `{"name":"{{flowName}}","status":"STARTED","defaultStatus":"STARTED"}`
}

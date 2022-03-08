package notifications

import (
	"github.com/aljrubior/go-mule/application"
	"strings"
)

const (
	FLOWS_PATTERN = "{{flows}}"
)

func NewPutDeploymentContextStartedNotification(serverId, contextId string, application *application.Application) *PutDeploymentContextStartedNotification {
	return &PutDeploymentContextStartedNotification{
		ApplicationNotification: ApplicationNotification{
			Notification: Notification{
				ServerId:  serverId,
				ContextId: contextId,
			},
			Application: application,
		},
	}
}

type PutDeploymentContextStartedNotification struct {
	ApplicationNotification
}

func (notification PutDeploymentContextStartedNotification) CreateNotification() string {
	message := notification.GetTemplate()

	message = strings.ReplaceAll(message, SERVER_ID_PATTERN, notification.ServerId)
	message = strings.ReplaceAll(message, CONTEXT_ID_PATTERN, notification.ContextId)
	message = strings.ReplaceAll(message, APPLICATION_NAME_PATTERN, notification.Application.Name)

	flows := notification.BuildFlows()
	message = strings.ReplaceAll(message, FLOWS_PATTERN, flows)

	return message
}

func (notification PutDeploymentContextStartedNotification) BuildFlows() string {
	var flows []string
	template := notification.GetFlowTemplate()

	for _, v := range notification.Application.FixedSchedulers {
		flows = append(flows, strings.ReplaceAll(template, FLOW_NAME_PATTERN, v.FlowName))
	}

	for _, v := range notification.Application.CronSchedulers {
		flows = append(flows, strings.ReplaceAll(template, FLOW_NAME_PATTERN, v.FlowName))
	}

	return strings.Join(flows, ",")
}

func (notification PutDeploymentContextStartedNotification) GetTemplate() string {
	return `PUT applications/{{applicationName}}/deployment HTTP/1.1
Content-Type: application/json
X-ANYPNT-SERVER-ID: {{serverId}}
X-ANYPNT-CTX-ID: {{contextId}}
X-ANYPNT-AGENT-VERSION: 2.4.28-SNAPSHOT
Content-Length: 40000
Content-Encoding: UTF-8

{"status":"CONTEXT_STARTED","message":"","application":{"name":"{{applicationName}}","domain":"default","state":"STARTED","flows":[{{flows}}],"lastDateStarted":1646568892451}}`
}

func (notification PutDeploymentContextStartedNotification) GetFlowTemplate() string {
	return `{"name":"{{flowName}}","status":"STARTED","defaultStatus":"STARTED"}`
}

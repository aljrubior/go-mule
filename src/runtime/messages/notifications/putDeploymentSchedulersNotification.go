package notifications

import (
	"github.com/aljrubior/go-mule/application"
	"strings"
)

const (
	SCHEDULER_PATTERN      = "{{schedulers}}"
	SCHEDULER_NAME_PATTHER = "{{schedulerName}}"
)

func NewPutDeploymentSchedulersNotification(serverId, contextId string, application *application.Application) *PutDeploymentSchedulersNotification {
	return &PutDeploymentSchedulersNotification{
		ApplicationNotification: ApplicationNotification{
			Notification: Notification{
				ServerId:  serverId,
				ContextId: contextId,
			},
			Application: application,
		},
	}
}

type PutDeploymentSchedulersNotification struct {
	ApplicationNotification
}

func (notification PutDeploymentSchedulersNotification) CreateNotification() string {
	message := notification.GetTemplate()

	message = strings.ReplaceAll(message, SERVER_ID_PATTERN, notification.ServerId)
	message = strings.ReplaceAll(message, CONTEXT_ID_PATTERN, notification.ContextId)
	message = strings.ReplaceAll(message, APPLICATION_NAME_PATTERN, notification.Application.Name)

	schedulers := notification.BuildSchedulers()
	message = strings.ReplaceAll(message, SCHEDULER_PATTERN, schedulers)

	return message
}

func (notification PutDeploymentSchedulersNotification) GetTemplate() string {
	return `PUT applications/{{applicationName}}/schedulers HTTP/1.1
Content-Type: application/json
X-ANYPNT-SERVER-ID: {{serverId}}
X-ANYPNT-CTX-ID: {{contextId}}
X-ANYPNT-AGENT-VERSION: 2.4.28-SNAPSHOT
Content-Length: 40000
Content-Encoding: UTF-8

{"applicationName":"{{applicationName}}","schedulers":[{{schedulers}}]}`
}

func (notification PutDeploymentSchedulersNotification) BuildSchedulers() string {
	var schedulers []string

	for _, v := range notification.Application.FixedSchedulers {
		template := notification.GetSchedulerTemplate(v.Name)
		template = strings.ReplaceAll(template, SCHEDULER_NAME_PATTHER, v.Name)
		template = strings.ReplaceAll(template, FLOW_NAME_PATTERN, v.FlowName)

		schedulers = append(schedulers, template)
	}

	for _, v := range notification.Application.CronSchedulers {
		template := notification.GetSchedulerTemplate(v.Name)
		template = strings.ReplaceAll(template, SCHEDULER_NAME_PATTHER, v.Name)
		template = strings.ReplaceAll(template, FLOW_NAME_PATTERN, v.FlowName)

		schedulers = append(schedulers, template)
	}

	return strings.Join(schedulers, ",")
}

func (notification PutDeploymentSchedulersNotification) GetSchedulerTemplate(flowName string) string {
	if strings.Contains(flowName, "Cron") {
		return notification.GetCronSchedulerTemplate()
	}

	return notification.GetFixedSchedulerTemplate()
}

func (notification PutDeploymentSchedulersNotification) GetCronSchedulerTemplate() string {
	return `{"type":"CronScheduler","name":"polling://{{flowName}}/","flowName":"{{flowName}}","enabled":true,"expression":"0 0/1 * 1/1 * ? *"`
}

func (notification PutDeploymentSchedulersNotification) GetFixedSchedulerTemplate() string {
	return `{"type":"FixedFrequencyScheduler","name":"{{schedulerName}}","flowName":"{{flowName}}","enabled":true,"timeUnit":"SECONDS","frequency":5,"startDelay":15}`
}

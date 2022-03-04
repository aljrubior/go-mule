package notifications

import "strings"

func NewPutDomainDeploymentDeployedNotification(serverId, contextId string) *PutDomainDeploymentDeployedNotification {
	return &PutDomainDeploymentDeployedNotification{
		serverId:  serverId,
		contextId: contextId,
	}
}

type PutDomainDeploymentDeployedNotification struct {
	serverId  string
	contextId string
	message   string
}

func (notification PutDomainDeploymentDeployedNotification) CreateNotification() string {
	message := `PUT domains/default/deployment HTTP/1.1
Content-Type: application/json
X-ANYPNT-SERVER-ID: {{serverId}}
X-ANYPNT-CTX-ID: {{contextId}}
X-ANYPNT-AGENT-VERSION: 2.4.28-SNAPSHOT
Content-Length: 80
Content-Encoding: UTF-8

{"status":"DEPLOYED","message":"","domain":{"name":"default","applications":[]}}`

	message = strings.Replace(message, SERVER_ID_PATTERN, notification.serverId, 1)
	message = strings.Replace(message, CONTEXT_ID_PATTERN, notification.contextId, 1)

	return message
}

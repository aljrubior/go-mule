package notifications

import "strings"

func NewPutDomainDeploymentContextInitialisedNotification(serverId, contextId string) *PutDomainDeploymentContextInitialisedNotification {
	return &PutDomainDeploymentContextInitialisedNotification{
		serverId:  serverId,
		contextId: contextId,
	}
}

type PutDomainDeploymentContextInitialisedNotification struct {
	serverId  string
	contextId string
	message   string
}

func (this PutDomainDeploymentContextInitialisedNotification) CreateNotification() string {
	message := `PUT domains/default/deployment HTTP/1.1
Content-Type: application/json
X-ANYPNT-SERVER-ID: {{serverId}}
X-ANYPNT-CTX-ID: {{contextId}}
X-ANYPNT-AGENT-VERSION: 2.4.28-SNAPSHOT
Content-Length: 91
Content-Encoding: UTF-8

{"status":"CONTEXT_INITIALISED","message":"","domain":{"name":"default","applications":[]}}`

	message = strings.Replace(message, SERVER_ID_PATTERN, this.serverId, 1)
	message = strings.Replace(message, CONTEXT_ID_PATTERN, this.contextId, 1)

	return message
}

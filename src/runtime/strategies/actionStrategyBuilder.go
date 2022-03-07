package strategies

import (
	"github.com/aljrubior/standalone-runtime/application"
	"github.com/aljrubior/standalone-runtime/runtime/messages"
	"github.com/gorilla/websocket"
)

func NewActionStrategyBuilder(
	conn *websocket.Conn,
	message *messages.WebsocketMessage,
	serverId,
	contextId string,
	applications *map[string]*application.Application,
	regex *messages.ActionRequestRegex) ActionStrategyBuilder {

	return ActionStrategyBuilder{
		conn:         conn,
		message:      message,
		serverId:     serverId,
		contextId:    contextId,
		applications: applications,
		regex:        regex,
	}
}

type ActionStrategyBuilder struct {
	conn    *websocket.Conn
	message *messages.WebsocketMessage
	serverId,
	contextId string
	applications *map[string]*application.Application
	regex        *messages.ActionRequestRegex
}

func (t ActionStrategyBuilder) Build() ActionStrategy {

	if t.regex.GetClusterAction.MatchString(t.message.GetResquestHeader()) {
		return NewGetClusterActionStrategy(t.conn, t.message)
	}

	if t.regex.PutStopApplication.MatchString(t.message.GetResquestHeader()) {
		applicationName := t.message.GetApplicationName()
		return NewPutStopApplicationActionStrategy(t.conn, t.message, t.serverId, t.contextId, (*t.applications)[applicationName])
	}

	if t.regex.PutStartApplication.MatchString(t.message.GetResquestHeader()) {
		applicationName := t.message.GetApplicationName()
		return NewPutStartApplicationActionStrategy(t.conn, t.message, t.serverId, t.contextId, (*t.applications)[applicationName])
	}

	if t.regex.PutApplication.MatchString(t.message.GetResquestHeader()) {
		applicationName := t.message.GetApplicationName()
		return NewPutApplicationActionStrategy(t.conn, t.message, t.serverId, t.contextId, (*t.applications)[applicationName])
	}

	if t.regex.GetAgentConfiguration.MatchString(t.message.GetResquestHeader()) {
		return NewGetAgentConfigurationActionStrategy(t.conn, t.message)
	}

	return NewNoActionStrategy()
}

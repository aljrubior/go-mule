package notifications

const (
	SERVER_ID_PATTERN  = "{{serverId}}"
	CONTEXT_ID_PATTERN = "{{contextId}}"
)

type Notification struct {
	ServerId  string
	ContextId string
}

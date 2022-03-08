package notifications

import "github.com/aljrubior/go-mule/application"

const (
	APPLICATION_NAME_PATTERN = "{{applicationName}}"
)

type ApplicationNotification struct {
	Notification
	Application *application.Application
}

package notifications

func NewPostKeepAliveNotification() *PostKeepAliveNotification {
	return &PostKeepAliveNotification{}
}

type PostKeepAliveNotification struct {
}

func (notification PostKeepAliveNotification) CreateNotification() string {
	return notification.GetTemplate()
}

func (notification PostKeepAliveNotification) GetTemplate() string {
	return `POST keepAlive HTTP/1.1`
}

package notifications

func NewPostKeepAliveNotification() *PostKeepAliveNotification {
	return &PostKeepAliveNotification{}
}

type PostKeepAliveNotification struct {
}

func (this PostKeepAliveNotification) CreateNotification() string {
	return this.GetTemplate()
}

func (this PostKeepAliveNotification) GetTemplate() string {
	return `POST keepAlive HTTP/1.1`
}

package handlers

type ServerHandler interface {
	CreateServer(token, serverName string) error
}

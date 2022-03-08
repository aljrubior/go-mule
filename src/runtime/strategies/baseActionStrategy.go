package strategies

import (
	"github.com/gorilla/websocket"
	"log"
	"time"
)

type BaseActionStrategy struct {
}

func (t *BaseActionStrategy) SendMessage(conn *websocket.Conn, notification string) {

	println(notification)

	if err := conn.WriteMessage(websocket.BinaryMessage, []byte(notification)); err != nil {
		log.Fatal(err)
	}

	time.Sleep(1 * time.Second)
}

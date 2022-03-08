package websockets

import (
	"crypto/tls"
	"github.com/gorilla/websocket"
	"net/http"
	"net/url"
	"time"
)

func NewRuntimeManagerDialer(url url.URL, config *tls.Config) *RuntimeManagerDialer {
	return &RuntimeManagerDialer{
		URL:       url,
		TLSConfig: config,
	}
}

type RuntimeManagerDialer struct {
	URL       url.URL
	TLSConfig *tls.Config
}

func (this RuntimeManagerDialer) CreateDialer() (*websocket.Conn, *http.Response, error) {

	dialer := websocket.Dialer{
		HandshakeTimeout: time.Second,
		ReadBufferSize:   1024,
		WriteBufferSize:  1024,
		TLSClientConfig:  this.TLSConfig,
	}

	return dialer.Dial(this.URL.String(), nil)
}

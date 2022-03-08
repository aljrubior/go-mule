package requests

import (
	"bytes"
	"fmt"
	"github.com/aljrubior/standalone-runtime/clients"
	"github.com/aljrubior/standalone-runtime/conf"
	"net/http"
)

func NewPostServerMetricsRequest(
	config *conf.MetricClientConfig,
	body []byte) *PostServerMetricsRequest {

	return &PostServerMetricsRequest{
		config: config,
		body:   body,
	}
}

type PostServerMetricsRequest struct {
	clients.BaseHttpRequest
	config *conf.MetricClientConfig
	body   []byte
}

func (t *PostServerMetricsRequest) buildUri() string {

	protocol := t.config.Protocol
	host := t.config.Host
	port := t.config.Port
	path := t.config.ServersPath

	return fmt.Sprintf("%s://%s:%s%s", protocol, host, port, path)
}

func (t *PostServerMetricsRequest) Build() *http.Request {

	uri := t.buildUri()

	req, _ := http.NewRequest(http.MethodPost, uri, bytes.NewBuffer(t.body))

	println(uri)

	req.Header.Add("Content-Type", clients.ContentTypeJSON)

	return req
}

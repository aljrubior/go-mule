package requests

import (
	"bytes"
	"fmt"
	"github.com/aljrubior/standalone-runtime/clients"
	"github.com/aljrubior/standalone-runtime/conf"
	"net/http"
)

func NewPostApplicationMetricsRequest(
	config *conf.MetricClientConfig,
	applicationName string,
	body []byte) *PostApplicationMetricsRequest {

	return &PostApplicationMetricsRequest{
		config:          config,
		applicationName: applicationName,
		body:            body,
	}
}

type PostApplicationMetricsRequest struct {
	clients.BaseHttpRequest
	config          *conf.MetricClientConfig
	applicationName string
	body            []byte
}

func (t *PostApplicationMetricsRequest) buildUri() string {

	protocol := t.config.Protocol
	host := t.config.Host
	port := t.config.Port
	path := t.config.ApplicationsPath

	return fmt.Sprintf("%s://%s:%s%s", protocol, host, port, path)
}

func (t *PostApplicationMetricsRequest) Build() *http.Request {

	uri := t.buildUri()

	req, _ := http.NewRequest(http.MethodPost, uri, bytes.NewBuffer(t.body))

	println(uri)

	req.Header.Add("Content-Type", clients.ContentTypeJSON)
	req.Header.Add("X-APPLICATION-NAME", t.applicationName)

	return req
}

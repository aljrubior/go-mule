package requests

import (
	"bytes"
	"fmt"
	"github.com/aljrubior/go-mule/clients"
	"github.com/aljrubior/go-mule/conf"
	"net/http"
)

func NewPostRegisterServerRequest(
	config *conf.ServerClientConfig,
	bearerToken,
	agentVersion string,
	body []byte) *PostRegisterServerRequest {

	return &PostRegisterServerRequest{
		config:       config,
		bearerToken:  bearerToken,
		agentVersion: agentVersion,
		body:         body,
	}
}

type PostRegisterServerRequest struct {
	clients.BaseHttpRequest
	config       *conf.ServerClientConfig
	bearerToken  string
	agentVersion string
	body         []byte
}

func (request *PostRegisterServerRequest) buildUri() string {

	protocol := request.config.Protocol
	host := request.config.Host
	path := request.config.ServersPath

	return fmt.Sprintf("%s://%s/%s", protocol, host, path)
}

func (request *PostRegisterServerRequest) Build() *http.Request {

	uri := request.buildUri()

	req, _ := http.NewRequest(http.MethodPost, uri, bytes.NewBuffer(request.body))

	req.Header.Add("Content-Type", clients.ContentTypeJSON)
	req.Header.Add("Authorization", request.GetBearerTokenValue(request.bearerToken))
	req.Header.Add("Accept-Language", "en-US,en;q=0.5")
	req.Header.Add("User-Agent", request.agentVersion)
	req.Header.Add("Accept", clients.ContentTypeJSON)

	return req
}

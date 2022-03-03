package requests

import (
	"bytes"
	"fmt"
	"github.com/aljrubior/standalone-runtime/clients"
	"github.com/aljrubior/standalone-runtime/conf"
	"net/http"
)

func NewPostRegisterServerRequest(
	config *conf.ServerClientConfig,
	bearerToken,
	organizationId,
	environmentId string,
	body []byte) *PostRegisterServerRequest {

	return &PostRegisterServerRequest{
		config:         config,
		bearerToken:    bearerToken,
		organizationId: organizationId,
		environmentId:  environmentId,
		body:           body,
	}
}

type PostRegisterServerRequest struct {
	clients.BaseHttpRequest
	config         *conf.ServerClientConfig
	bearerToken    string
	organizationId string
	environmentId  string
	body           []byte
}

func (this *PostRegisterServerRequest) buildUri() string {

	protocol := this.config.Protocol
	host := this.config.Host
	path := this.config.ServersPath

	return fmt.Sprintf("%s://%s/%s", protocol, host, path)
}

func (this *PostRegisterServerRequest) Build() *http.Request {

	uri := this.buildUri()

	req, _ := http.NewRequest(http.MethodPost, uri, bytes.NewBuffer(this.body))

	this.AddDefaultHeaders(req, this.organizationId, this.environmentId, this.bearerToken)
	this.AddContentType(req, clients.ContentTypeJSON)

	return req
}

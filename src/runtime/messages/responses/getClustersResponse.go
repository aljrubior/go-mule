package responses

import "strings"

func NewGetClustersResponse(messageId string) *GetClustersResponse {
	return &GetClustersResponse{
		Response{
			MessageId: messageId,
		},
	}
}

type GetClustersResponse struct {
	Response
}

func (this GetClustersResponse) CreateResponse() string {

	message := this.GetTemplate()

	return strings.Replace(message, MESSAGE_ID_PATTERN, this.MessageId, 1)
}

func (this GetClustersResponse) GetTemplate() string {

	return `HTTP/1.1 404 Not Found
Message-Id: {{messageId}}
Content-Type: application/json
Content-Length: 89

{"errorType":"ServerNotClustered", "errorMessage":"The server is not part of a cluster" }`
}

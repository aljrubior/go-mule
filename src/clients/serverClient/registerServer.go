package serverClient

import (
	"encoding/json"
	"github.com/aljrubior/go-mule/clients/serverClient/requests"
	"github.com/aljrubior/go-mule/clients/serverClient/responses"
	"io/ioutil"
	"net/http"
	"time"
)

func (serverClient DefaultServerClient) RegisterServer(token, agentVersion string, body []byte) (*responses.ServerRegistrationResponse, error) {

	client := &http.Client{Timeout: time.Duration(10) * time.Second}

	req := requests.NewPostRegisterServerRequest(&serverClient.config, token, agentVersion, body).Build()

	resp, err := client.Do(req)

	defer resp.Body.Close()

	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, serverClient.ThrowError(resp)
	}

	data, err := ioutil.ReadAll(resp.Body)

	println(string(data))

	if err != nil {
		return nil, err
	}

	var response responses.ServerRegistrationResponse

	err = json.Unmarshal(data, &response)

	if err != nil {
		return nil, err
	}

	return &response, nil
}

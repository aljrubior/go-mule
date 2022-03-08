package metricClient

import (
	"github.com/aljrubior/standalone-runtime/clients/metricClient/requests"
	"io/ioutil"
	"net/http"
	"time"
)

func (t *DefaultMetricClient) PostServerMetrics(body []byte) error {

	transport := &http.Transport{TLSClientConfig: t.tlsConfig}

	client := &http.Client{
		Transport: transport,
		Timeout:   time.Duration(10) * time.Second,
	}

	req := requests.NewPostServerMetricsRequest(t.config, body).Build()

	resp, err := client.Do(req)

	defer resp.Body.Close()

	if err != nil {
		return err
	}

	println("Status Code:", resp.StatusCode)

	if resp.StatusCode != 200 {
		return t.ThrowError(resp)
	}

	data, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return err
	}

	println("Response:", string(data))

	return nil
}

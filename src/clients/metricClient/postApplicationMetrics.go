package metricClient

import (
	"github.com/aljrubior/go-mule/clients/metricClient/requests"
	"io/ioutil"
	"net/http"
	"time"
)

func (t *DefaultMetricClient) PostApplicationMetrics(applicationName string, body []byte) error {

	transport := &http.Transport{TLSClientConfig: t.tlsConfig}

	client := &http.Client{
		Transport: transport,
		Timeout:   time.Duration(10) * time.Second,
	}

	req := requests.NewPostApplicationMetricsRequest(t.config, applicationName, body).Build()

	resp, err := client.Do(req)

	defer resp.Body.Close()

	if err != nil {
		return err
	}

	if resp.StatusCode != 204 {
		return t.ThrowError(resp)
	}

	_, err = ioutil.ReadAll(resp.Body)

	return err
}

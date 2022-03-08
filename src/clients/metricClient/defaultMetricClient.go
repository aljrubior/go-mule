package metricClient

import (
	"crypto/tls"
	"github.com/aljrubior/standalone-runtime/clients"
	"github.com/aljrubior/standalone-runtime/conf"
)

func NewDefaultMetricClient(config *conf.MetricClientConfig, tlsConfig *tls.Config) DefaultMetricClient {
	return DefaultMetricClient{
		config:    config,
		tlsConfig: tlsConfig,
	}
}

type DefaultMetricClient struct {
	clients.BaseHttpClient
	config    *conf.MetricClientConfig
	tlsConfig *tls.Config
}

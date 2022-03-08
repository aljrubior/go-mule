package messages

import "regexp"

const (
	MESSAGE_ID_HEADER                          = "Message-Id"
	GET_CLUSTERS_REQUEST_PATTERN               = "GET clusters HTTP/1.1"
	GET_AGENT_CONFIGURATION_REQUEST_PATTERN    = "GET agent/configuration HTTP/1.1"
	PUT_APPLICATIONS_REQUEST_PATTERN           = `PUT applications/([a-z0-9\-]+) HTTP/1.1`
	PUT_APPLICATIONS_STOP_REQUEST_PATTERN      = `PUT applications/([a-z0-9\-]+)/stop HTTP/1.1`
	PUT_APPLICATIONS_START_REQUEST_PATTERN     = `PUT applications/([a-z0-9\-]+)/start HTTP/1.1`
	PATCH_LOGGING_SERVICE_REQUEST_PATTERN      = "PATCH agent/mule.agent.logging.service HTTP/1.1"
	PUT_LOGGING_SERVICE_ENABLE_REQUEST_PATTERN = "PUT agent/mule.agent.logging.service/enable HTTP/1.1"
	PUT_TRACKING_SERVICE_REQUEST_PATTERN       = "PUT agent/mule.agent.tracking.service HTTP/1.1"
	PUT_TRACKING_SERVICE_ENABLE_REQUEST_ACTION = "PUT agent/mule.agent.tracking.service/enable HTTP/1.1"
)

func NewActionRequestRegex() ActionRequestRegex {

	regex := ActionRequestRegex{}

	regex.init()

	return regex
}

type ActionRequestRegex struct {
	GetClusterAction,
	GetAgentConfiguration,
	PutApplication,
	PutStopApplication,
	PutStartApplication,
	PatchLoggingService,
	PutEnableLoggingService,
	PutTrackingService,
	PutEnableTrackingService *regexp.Regexp
}

func (t *ActionRequestRegex) init() {

	t.GetClusterAction, _ = regexp.Compile(GET_CLUSTERS_REQUEST_PATTERN)
	t.GetAgentConfiguration, _ = regexp.Compile(GET_AGENT_CONFIGURATION_REQUEST_PATTERN)
	t.PutApplication, _ = regexp.Compile(PUT_APPLICATIONS_REQUEST_PATTERN)
	t.PutStopApplication, _ = regexp.Compile(PUT_APPLICATIONS_STOP_REQUEST_PATTERN)
	t.PutStartApplication, _ = regexp.Compile(PUT_APPLICATIONS_START_REQUEST_PATTERN)
	t.PatchLoggingService, _ = regexp.Compile(PATCH_LOGGING_SERVICE_REQUEST_PATTERN)
	t.PutEnableLoggingService, _ = regexp.Compile(PUT_LOGGING_SERVICE_ENABLE_REQUEST_PATTERN)
	t.PutTrackingService, _ = regexp.Compile(PUT_TRACKING_SERVICE_REQUEST_PATTERN)
	t.PutEnableTrackingService, _ = regexp.Compile(PUT_TRACKING_SERVICE_ENABLE_REQUEST_ACTION)
}

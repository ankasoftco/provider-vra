package models 
type HealthCheckConfiguration struct {

	// Number of consecutive successful checks before considering a particular back-end instance as healthy.
	// Example: 2
	HealthyThreshold int32 `json:"healthyThreshold,omitempty"`

	// HTTP or HTTPS method to use when sending a health check request.
	// Example: GET, OPTIONS, POST, HEAD, PUT
	HTTPMethod string `json:"httpMethod,omitempty"`

	// Interval (in seconds) at which the health checks will be performed.
	// Example: 60
	IntervalSeconds int32 `json:"intervalSeconds,omitempty"`

	// Enable passive monitor mode. This setting only applies to NSX-T.
	// Example: false
	PassiveMonitor bool `json:"passiveMonitor,omitempty"`

	// Port on the back-end instance machine to use for the health check.
	// Example: 80
	Port string `json:"port,omitempty"`

	// The protocol used for the health check.
	// Example: HTTP, HTTPS
	Protocol string `json:"protocol,omitempty"`

	// Request body. Used by HTTP, HTTPS, TCP, UDP.
	// Example: http_request.body
	RequestBody string `json:"requestBody,omitempty"`

	// Expected response body. Used by HTTP, HTTPS, TCP, UDP.
	// Example: http_response.body
	ResponseBody string `json:"responseBody,omitempty"`

	// Timeout (in seconds) to wait for a response from the back-end instance.
	// Example: 5
	TimeoutSeconds int32 `json:"timeoutSeconds,omitempty"`

	// Number of consecutive check failures before considering a particular back-end instance as unhealthy.
	// Example: 5
	UnhealthyThreshold int32 `json:"unhealthyThreshold,omitempty"`

	// URL path on the back-end instance against which a request will be performed for the health check. Useful when the health check protocol is HTTP/HTTPS.
	// Example: /index.html
	URLPath string `json:"urlPath,omitempty"`
}


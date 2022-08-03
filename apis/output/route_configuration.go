package models 
type RouteConfiguration struct {

	// Algorithm employed for load balancing.
	// Example: ROUND_ROBIN
	Algorithm string `json:"algorithm,omitempty"`

	// Parameters need for load balancing algorithm.Use newline to separate multiple parameters.
	// Example: uriLength=10\nurlParam=section
	AlgorithmParameters string `json:"algorithmParameters,omitempty"`

	// Health check configuration for this route configuration.
	HealthCheckConfiguration *HealthCheckConfiguration `json:"healthCheckConfiguration,omitempty"`

	// Member port where the traffic is routed to.
	// Example: 80
	// Required: true
	MemberPort *string `json:"memberPort"`

	// The protocol of the member traffic.
	// Example: TCP, UDP
	// Required: true
	MemberProtocol *string `json:"memberProtocol"`

	// Port which the load balancer is listening to.
	// Example: 80
	// Required: true
	Port *string `json:"port"`

	// The protocol of the incoming load balancer requests.
	// Example: TCP, UDP
	// Required: true
	Protocol *string `json:"protocol"`
}


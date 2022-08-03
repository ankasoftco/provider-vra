package models 
type LoadBalancerSpecification struct {

	// Additional custom properties that may be used to extend this resource.
	CustomProperties map[string]string `json:"customProperties,omitempty"`

	// The id of the deployment that is associated with this resource
	// Example: 123e4567-e89b-12d3-a456-426655440000
	DeploymentID string `json:"deploymentId,omitempty"`

	// A human-friendly description.
	Description string `json:"description,omitempty"`

	// An Internet-facing load balancer has a publicly resolvable DNS name, so it can route requests from clients over the Internet to the instances that are registered with the load balancer.
	InternetFacing bool `json:"internetFacing,omitempty"`

	// Defines logging level for collecting load balancer traffic logs.
	// Example: ERROR, WARNING, INFO, DEBUG
	LoggingLevel string `json:"loggingLevel,omitempty"`

	// A human-friendly name used as an identifier in APIs that support this option.
	// Required: true
	Name *string `json:"name"`

	// A set of network interface specifications for this load balancer.
	// Required: true
	Nics []*NetworkInterfaceSpecification `json:"nics"`

	// The id of the project the current user belongs to.
	// Example: e058
	// Required: true
	ProjectID *string `json:"projectId"`

	// The load balancer route configuration regarding ports and protocols.
	// Required: true
	Routes []*RouteConfiguration `json:"routes"`

	// A set of tag keys and optional values that should be set on any resource that is produced from this specification.
	// Example: [ { \"key\" : \"ownedBy\", \"value\": \"Rainpole\" } ]
	Tags []*Tag `json:"tags"`

	// A list of links to target load balancer pool members. Links can be to either a machine or a machine's network interface.
	// Example: [ \"/iaas/machines/eac3d\" ]
	TargetLinks []string `json:"targetLinks"`

	// Define the type/variant of load balancer numbers e.g.for NSX the number virtual servers and pool members load balancer can host
	// Example: SMALL, MEDIUM, LARGE
	Type string `json:"type,omitempty"`
}


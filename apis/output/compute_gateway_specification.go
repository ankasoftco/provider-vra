package models 
type ComputeGatewaySpecification struct {

	// Additional custom properties that may be used to extend this resource.
	CustomProperties map[string]string `json:"customProperties,omitempty"`

	// The id of the deployment that is associated with this resource
	// Example: 123e4567-e89b-12d3-a456-426655440000
	DeploymentID string `json:"deploymentId,omitempty"`

	// A human-friendly name used as an identifier in APIs that support this option.
	// Required: true
	Name *string `json:"name"`

	// List of NAT Rules
	// Required: true
	NatRules []*NatRule `json:"natRules"`

	// List of networks
	// Required: true
	Networks []string `json:"networks"`

	// The id of the project the current user belongs to.
	// Example: e058
	// Required: true
	ProjectID *string `json:"projectId"`
}


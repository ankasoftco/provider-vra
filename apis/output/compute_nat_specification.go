package models 
type ComputeNatSpecification struct {

	// Additional custom properties that may be used to extend this resource.
	CustomProperties map[string]string `json:"customProperties,omitempty"`

	// The id of the deployment that is associated with this resource
	// Example: 123e4567-e89b-12d3-a456-426655440000
	DeploymentID string `json:"deploymentId,omitempty"`

	// Id of the Compute Gateway to which the Compute Nat resource will be attached
	// Required: true
	Gateway *string `json:"gateway"`

	// A human-friendly name used as an identifier in APIs that support this option.
	// Required: true
	Name *string `json:"name"`

	// List of NAT Rules
	// Required: true
	NatRules []*NatRule `json:"natRules"`

	// The id of the project the current user belongs to.
	// Example: e058
	// Required: true
	ProjectID *string `json:"projectId"`
}


package models 
type ResourceSpecification struct {

	// Resource deployment id
	DeploymentID string `json:"deploymentId,omitempty"`

	// Name of the resource
	// Required: true
	Name *string `json:"name"`

	// Resource project id
	ProjectID string `json:"projectId,omitempty"`

	// Resource properties
	Properties interface{} `json:"properties,omitempty"`

	// Type of the resource
	// Required: true
	Type *string `json:"type"`
}


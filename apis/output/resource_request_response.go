package models 
type ResourceRequestResponse struct {

	// Identifier of the requested deployment id to which the request applies to
	// Format: uuid
	DeploymentID strfmt.UUID `json:"deploymentId,omitempty"`

	// Project identifier
	ProjectID string `json:"projectId,omitempty"`

	// Request identifier
	// Format: uuid
	RequestID strfmt.UUID `json:"requestId,omitempty"`

	// Resource ID
	ResourceID string `json:"resourceId,omitempty"`
}


package models 
type CatalogItemRequestResponse struct {

	// The created deployment's ID
	DeploymentID string `json:"deploymentId,omitempty"`

	// The created deployment's name
	DeploymentName string `json:"deploymentName,omitempty"`
}


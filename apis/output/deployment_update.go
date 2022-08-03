package models 
type DeploymentUpdate struct {

	// New description of the deployment
	Description string `json:"description,omitempty"`

	// New iconid of the deployment
	// Format: uuid
	IconID strfmt.UUID `json:"iconId,omitempty"`

	// New name of the deployment
	Name string `json:"name,omitempty"`
}


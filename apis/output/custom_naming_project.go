package models 
type CustomNamingProject struct {

	// Flag to check if project is active
	// Example: true
	Active bool `json:"active,omitempty"`

	// Flag to represent if custom name is default for org
	// Example: true
	DefaultOrg bool `json:"defaultOrg,omitempty"`

	// Unique id of custom naming project
	// Example: 3fa85f64-5717-4562-b3fc-2c963f66afa6
	// Format: uuid
	ID strfmt.UUID `json:"id,omitempty"`

	// Org id
	// Example: 3fa85f64-5717-4562-b3fc-2c963f76afa6
	OrgID string `json:"orgId,omitempty"`

	// Project id mapped to custom name
	// Example: 3fa85f64-5717-4562-b3fc-2c963f66afa7
	ProjectID string `json:"projectId,omitempty"`

	// Name of mapped project
	// Example: Project name
	ProjectName string `json:"projectName,omitempty"`
}


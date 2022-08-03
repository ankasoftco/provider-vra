package models 
type CnProjectVo struct {

	// active
	Active bool `json:"active,omitempty"`

	// default org
	DefaultOrg bool `json:"defaultOrg,omitempty"`

	// id
	// Format: uuid
	ID strfmt.UUID `json:"id,omitempty"`

	// org Id
	OrgID string `json:"orgId,omitempty"`

	// project Id
	ProjectID string `json:"projectId,omitempty"`

	// project name
	ProjectName string `json:"projectName,omitempty"`
}


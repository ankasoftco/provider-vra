package models 
type CnTemplateCounterVo struct {

	// active
	Active bool `json:"active,omitempty"`

	// cn resource type
	// Enum: [COMPUTE NETWORK COMPUTE_STORAGE LOAD_BALANCER RESOURCE_GROUP GATEWAY NAT SECURITY_GROUP]
	CnResourceType string `json:"cnResourceType,omitempty"`

	// current counter
	CurrentCounter int64 `json:"currentCounter,omitempty"`

	// id
	// Format: uuid
	ID strfmt.UUID `json:"id,omitempty"`

	// project Id
	ProjectID string `json:"projectId,omitempty"`
}


package models 
type CustomNamingCounter struct {

	// active
	Active bool `json:"active,omitempty"`

	// The resource type of custom name
	// Required: true
	// Enum: [COMPUTE NETWORK COMPUTE_STORAGE LOAD_BALANCER RESOURCE_GROUP GATEWAY NAT SECURITY_GROUP]
	CnResourceType *string `json:"cnResourceType"`

	// The current counter of custom name
	// Required: true
	CurrentCounter *int64 `json:"currentCounter"`

	// id
	// Format: uuid
	ID strfmt.UUID `json:"id,omitempty"`

	// The project id to which the counter is mapped
	// Required: true
	ProjectID *string `json:"projectId"`
}


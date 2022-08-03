package models 
type CnTemplateVo struct {

	// counters
	Counters []*CnTemplateCounterVo `json:"counters"`

	// id
	// Format: uuid
	ID strfmt.UUID `json:"id,omitempty"`

	// increment step
	IncrementStep int64 `json:"incrementStep,omitempty"`

	// pattern
	Pattern string `json:"pattern,omitempty"`

	// resource default
	ResourceDefault bool `json:"resourceDefault,omitempty"`

	// resource type
	// Enum: [COMPUTE NETWORK COMPUTE_STORAGE LOAD_BALANCER RESOURCE_GROUP GATEWAY NAT SECURITY_GROUP]
	ResourceType string `json:"resourceType,omitempty"`

	// resource type name
	ResourceTypeName string `json:"resourceTypeName,omitempty"`

	// start counter
	StartCounter int64 `json:"startCounter,omitempty"`

	// static pattern
	StaticPattern string `json:"staticPattern,omitempty"`

	// unique name
	UniqueName bool `json:"uniqueName,omitempty"`
}


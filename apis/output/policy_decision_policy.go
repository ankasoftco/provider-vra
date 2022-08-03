package models 
type PolicyDecisionPolicy struct {

	// enforcement type
	// Enum: [SOFT HARD]
	EnforcementType string `json:"enforcementType,omitempty"`

	// id
	// Format: uuid
	ID strfmt.UUID `json:"id,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// project Id
	ProjectID string `json:"projectId,omitempty"`

	// rank
	Rank int32 `json:"rank,omitempty"`

	// status
	// Enum: [NOT_ENFORCED ENFORCED CONFLICT]
	Status string `json:"status,omitempty"`
}


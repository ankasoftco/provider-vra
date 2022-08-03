package models 
type DeploymentResource struct {

	// Creation time
	// Format: date-time
	CreatedAt strfmt.DateTime `json:"createdAt,omitempty"`

	// Current request
	CurrentRequest *Request `json:"currentRequest,omitempty"`

	// A list of other resources this resource depends on
	DependsOn []string `json:"dependsOn"`

	// A description of the resource
	Description string `json:"description,omitempty"`

	// Expense associated with the deployment.
	// Read Only: true
	Expense *Expense `json:"expense,omitempty"`

	// Unique identifier of the resource
	// Format: uuid
	ID strfmt.UUID `json:"id,omitempty"`

	// Name of the resource
	// Required: true
	Name *string `json:"name"`

	// Origin of the resource
	// Enum: [DISCOVERED ONBOARDED MIGRATED]
	Origin string `json:"origin,omitempty"`

	// properties
	Properties interface{} `json:"properties,omitempty"`

	// The current state of the resource
	// Enum: [PARTIAL TAINTED OK]
	State string `json:"state,omitempty"`

	// The current sync status
	// Enum: [SUCCESS MISSING STALE]
	SyncStatus string `json:"syncStatus,omitempty"`

	// Type of the resource
	// Required: true
	Type *string `json:"type"`
}


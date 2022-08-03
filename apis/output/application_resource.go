package models 
type ApplicationResource struct {

	// Creation time
	// Format: date-time
	CreatedAt strfmt.DateTime `json:"createdAt,omitempty"`

	// A description of the resource
	Description string `json:"description,omitempty"`

	// Expense associated with the deployment.
	// Read Only: true
	Expense *Expense `json:"expense,omitempty"`

	// Unique identifier of the resource
	ID string `json:"id,omitempty"`

	// Name of the resource
	// Required: true
	Name *string `json:"name"`

	// Origin of the resource
	// Enum: [DISCOVERED ONBOARDED MIGRATED]
	Origin string `json:"origin,omitempty"`

	// properties
	Properties interface{} `json:"properties,omitempty"`

	// The current sync status
	// Enum: [SUCCESS MISSING STALE]
	SyncStatus string `json:"syncStatus,omitempty"`

	// tier
	Tier *ApplicationTier `json:"tier,omitempty"`

	// Type of the resource
	// Required: true
	Type *string `json:"type"`
}


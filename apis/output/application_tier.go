package models 
type ApplicationTier struct {

	// A description of the tier
	Description string `json:"description,omitempty"`

	// Unique identifier of the tier
	ID string `json:"id,omitempty"`

	// Name of the tier
	// Required: true
	Name *string `json:"name"`
}


package models 
type Application struct {

	// Creation time
	// Format: date-time
	CreatedAt strfmt.DateTime `json:"createdAt,omitempty"`

	// Description of the application
	Description string `json:"description,omitempty"`

	// Id of the application
	ID string `json:"id,omitempty"`

	// Last update time
	// Format: date-time
	LastUpdatedAt strfmt.DateTime `json:"lastUpdatedAt,omitempty"`

	// Name of the application
	// Required: true
	Name *string `json:"name"`

	// org Id
	OrgID string `json:"orgId,omitempty"`

	// Resources of the application.
	Resources []*ApplicationResource `json:"resources"`

	// Tiers of the application.
	Tiers []*ApplicationTier `json:"tiers"`
}


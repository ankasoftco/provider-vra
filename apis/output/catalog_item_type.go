package models 
type CatalogItemType struct {

	// Base callback URI for import operations
	BaseURI string `json:"baseUri,omitempty"`

	// Configuration schema for this type, requests to create a source of this type should have a compliant 'config' field
	ConfigSchema interface{} `json:"configSchema,omitempty"`

	// Creation time
	// Format: date-time
	CreatedAt strfmt.DateTime `json:"createdAt,omitempty"`

	// Service that created this type
	CreatedBy string `json:"createdBy,omitempty"`

	// Default Icon Id
	// Format: uuid
	IconID strfmt.UUID `json:"iconId,omitempty"`

	// Human-readable unique ID containing only lowercase letters and periods, neither starting nor ending with a period, and never having two consecutive periods
	ID string `json:"id,omitempty"`

	// Unique Catalog Item Type name
	// Required: true
	Name *string `json:"name"`
}


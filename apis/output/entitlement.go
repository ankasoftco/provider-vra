package models 
type Entitlement struct {

	// Entitlement definition that contains the Catalog Item or Catalog Source data.
	// Required: true
	Definition *ContentDefinition `json:"definition"`

	// Entitlement id
	// Required: true
	// Format: uuid
	ID *strfmt.UUID `json:"id"`

	// Project id
	// Required: true
	ProjectID *string `json:"projectId"`
}


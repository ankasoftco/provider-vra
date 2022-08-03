package models 
type CatalogSource struct {

	// Source custom configuration
	// Required: true
	Config interface{} `json:"config"`

	// Creation time
	// Format: date-time
	CreatedAt strfmt.DateTime `json:"createdAt,omitempty"`

	// Created By
	CreatedBy string `json:"createdBy,omitempty"`

	// Catalog Source description
	Description string `json:"description,omitempty"`

	// Global flag indicating that all the items can be requested across all projects.
	Global bool `json:"global,omitempty"`

	// Default Icon Id
	// Format: uuid
	IconID strfmt.UUID `json:"iconId,omitempty"`

	// Catalog Source id
	// Required: true
	// Format: uuid
	ID *strfmt.UUID `json:"id"`

	// Number of items found
	ItemsFound int32 `json:"itemsFound,omitempty"`

	// Number of items imported.
	ItemsImported int32 `json:"itemsImported,omitempty"`

	// Last import completion time
	// Format: date-time
	LastImportCompletedAt strfmt.DateTime `json:"lastImportCompletedAt,omitempty"`

	// Last import error(s)
	LastImportErrors []string `json:"lastImportErrors"`

	// Last import start time
	// Format: date-time
	LastImportStartedAt strfmt.DateTime `json:"lastImportStartedAt,omitempty"`

	// Update time
	// Format: date-time
	LastUpdatedAt strfmt.DateTime `json:"lastUpdatedAt,omitempty"`

	// Updated By
	LastUpdatedBy string `json:"lastUpdatedBy,omitempty"`

	// Catalog Source name
	// Required: true
	Name *string `json:"name"`

	// Project id where the source belongs
	ProjectID string `json:"projectId,omitempty"`

	// Type of source, e.g. blueprint, CFT... etc
	// Required: true
	TypeID *string `json:"typeId"`
}


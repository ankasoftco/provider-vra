package models 
type CatalogItem struct {

	// Max number of instances that can be requested at a time
	// Maximum: 127
	// Minimum: -128
	BulkRequestLimit *int32 `json:"bulkRequestLimit,omitempty"`

	// Creation time
	// Format: date-time
	CreatedAt strfmt.DateTime `json:"createdAt,omitempty"`

	// Created By
	CreatedBy string `json:"createdBy,omitempty"`

	// CatalogItem description
	Description string `json:"description,omitempty"`

	// Form ID
	FormID string `json:"formId,omitempty"`

	// Icon ID
	// Format: uuid
	IconID strfmt.UUID `json:"iconId,omitempty"`

	// CatalogItem id
	// Required: true
	// Format: uuid
	ID *strfmt.UUID `json:"id"`

	// Update time
	// Format: date-time
	LastUpdatedAt strfmt.DateTime `json:"lastUpdatedAt,omitempty"`

	// Updated By
	LastUpdatedBy string `json:"lastUpdatedBy,omitempty"`

	// CatalogItem name
	// Required: true
	Name *string `json:"name"`

	// Associated project IDs that can be used for requesting
	// Required: true
	ProjectIds []string `json:"projectIds"`

	// Associated projects that can be used for requesting
	Projects []*ResourceReference `json:"projects"`

	// Json schema describing request parameters, a simplified version of http://json-schema.org/latest/json-schema-validation.html#rfc.section.5
	Schema interface{} `json:"schema,omitempty"`

	// LibraryItem source ID.
	// Format: uuid
	SourceID strfmt.UUID `json:"sourceId,omitempty"`

	// LibraryItem source name.
	SourceName string `json:"sourceName,omitempty"`

	// ResourceReference to type, e.g. blueprint, CFT... etc
	// Required: true
	Type *ResourceReference `json:"type"`
}


package models 
type CatalogItemVersion struct {

	// Created At
	// Format: date-time
	CreatedAt strfmt.DateTime `json:"createdAt,omitempty"`

	// Description
	Description string `json:"description,omitempty"`

	// Form ID
	FormID string `json:"formId,omitempty"`

	// Version ID
	ID string `json:"id,omitempty"`

	// Input Schema
	Schema interface{} `json:"schema,omitempty"`
}


package models 
type PropertyGroup struct {

	// Created time
	// Read Only: true
	// Format: date-time
	CreatedAt strfmt.DateTime `json:"createdAt,omitempty"`

	// Created by
	// Read Only: true
	CreatedBy string `json:"createdBy,omitempty"`

	// Property group description
	Description string `json:"description,omitempty"`

	// Property group display name
	DisplayName string `json:"displayName,omitempty"`

	// Object ID
	// Read Only: true
	ID string `json:"id,omitempty"`

	// Property group name
	Name string `json:"name,omitempty"`

	// Org ID
	// Read Only: true
	OrgID string `json:"orgId,omitempty"`

	// Project ID
	ProjectID string `json:"projectId,omitempty"`

	// Project Name
	// Read Only: true
	ProjectName string `json:"projectName,omitempty"`

	// Properties
	Properties map[string]Property `json:"properties,omitempty"`

	// Property group type
	// Enum: [INPUT CONSTANT]
	Type string `json:"type,omitempty"`

	// Updated time
	// Read Only: true
	// Format: date-time
	UpdatedAt strfmt.DateTime `json:"updatedAt,omitempty"`

	// Updated by
	// Read Only: true
	UpdatedBy string `json:"updatedBy,omitempty"`
}


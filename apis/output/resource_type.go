package models 
type ResourceType struct {

	// Created time
	// Read Only: true
	// Format: date-time
	CreatedAt strfmt.DateTime `json:"createdAt,omitempty"`

	// Created by
	// Read Only: true
	CreatedBy string `json:"createdBy,omitempty"`

	// Resource type description
	// Read Only: true
	Description string `json:"description,omitempty"`

	// Resource type display name
	// Read Only: true
	DisplayName string `json:"displayName,omitempty"`

	// Endpoint type
	// Read Only: true
	EndpointType string `json:"endpointType,omitempty"`

	// Object ID
	// Read Only: true
	ID string `json:"id,omitempty"`

	// Resource type name
	// Read Only: true
	Name string `json:"name,omitempty"`

	// Provider resource operations
	// Read Only: true
	Operations map[string]string `json:"operations,omitempty"`

	// Org ID
	// Read Only: true
	OrgID string `json:"orgId,omitempty"`

	// Provider Id
	// Read Only: true
	ProviderID string `json:"providerId,omitempty"`

	// Provider name
	// Read Only: true
	ProviderName string `json:"providerName,omitempty"`

	// Provider version
	// Read Only: true
	ProviderVersion string `json:"providerVersion,omitempty"`

	// Resource type schema
	// Read Only: true
	Schema *PropertyDefinition `json:"schema,omitempty"`

	// Updated time
	// Read Only: true
	// Format: date-time
	UpdatedAt strfmt.DateTime `json:"updatedAt,omitempty"`

	// Updated by
	// Read Only: true
	UpdatedBy string `json:"updatedBy,omitempty"`
}


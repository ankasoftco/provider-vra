package models 
type ContentSource struct {

	// Source custom configuration
	// Example: {"branch":"string","contentType":"string","endpointId":"string","integrationId":"string","path":"string","repository":"string"}
	Config interface{} `json:"config,omitempty"`

	// Creation time
	// Read Only: true
	// Format: date-time
	CreatedAt strfmt.DateTime `json:"createdAt,omitempty"`

	// Created By
	// Read Only: true
	CreatedBy string `json:"createdBy,omitempty"`

	// Content Source description
	Description string `json:"description,omitempty"`

	// Content Source id
	// Required: true
	// Format: uuid
	ID *strfmt.UUID `json:"id"`

	// Update time
	// Read Only: true
	// Format: date-time
	LastUpdatedAt strfmt.DateTime `json:"lastUpdatedAt,omitempty"`

	// Updated By
	// Read Only: true
	LastUpdatedBy string `json:"lastUpdatedBy,omitempty"`

	// Content Source name
	// Required: true
	Name *string `json:"name"`

	// Associated org
	// Read Only: true
	OrgID string `json:"orgId,omitempty"`

	// Associated projects
	// Required: true
	ProjectID *string `json:"projectId"`

	// Is Sync Enabled
	// Example: false
	SyncEnabled bool `json:"syncEnabled,omitempty"`

	// Content Source type
	// Required: true
	// Enum: [com.github com.gitlab org.bitbucket com.vmware.marketplace]
	TypeID *string `json:"typeId"`
}


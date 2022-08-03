package models 
type BlueprintVersion struct {

	// Draft blueprint ID
	// Read Only: true
	BlueprintID string `json:"blueprintId,omitempty"`

	// Blueprint YAML content
	// Read Only: true
	Content string `json:"content,omitempty"`

	// Created time
	// Read Only: true
	// Format: date-time
	CreatedAt strfmt.DateTime `json:"createdAt,omitempty"`

	// Created by
	// Read Only: true
	CreatedBy string `json:"createdBy,omitempty"`

	// Draft blueprint description
	// Read Only: true
	Description string `json:"description,omitempty"`

	// Object ID
	// Read Only: true
	ID string `json:"id,omitempty"`

	// Blueprint name
	// Read Only: true
	Name string `json:"name,omitempty"`

	// Org ID
	// Read Only: true
	OrgID string `json:"orgId,omitempty"`

	// Project ID
	ProjectID string `json:"projectId,omitempty"`

	// Project Name
	// Read Only: true
	ProjectName string `json:"projectName,omitempty"`

	// Blueprint version self link
	// Read Only: true
	SelfLink string `json:"selfLink,omitempty"`

	// Blueprint status
	// Read Only: true
	// Enum: [DRAFT VERSIONED RELEASED]
	Status string `json:"status,omitempty"`

	// Updated time
	// Read Only: true
	// Format: date-time
	UpdatedAt strfmt.DateTime `json:"updatedAt,omitempty"`

	// Updated by
	// Read Only: true
	UpdatedBy string `json:"updatedBy,omitempty"`

	// Validation result
	// Read Only: true
	Valid *bool `json:"valid,omitempty"`

	// Blueprint version
	// Read Only: true
	Version string `json:"version,omitempty"`

	// Blueprint version change log
	// Read Only: true
	VersionChangeLog string `json:"versionChangeLog,omitempty"`

	// Blueprint version description
	// Read Only: true
	VersionDescription string `json:"versionDescription,omitempty"`
}


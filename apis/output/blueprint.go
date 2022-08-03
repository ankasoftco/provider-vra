package models 
type Blueprint struct {

	// Blueprint YAML content
	Content string `json:"content,omitempty"`

	// Content source id
	// Read Only: true
	ContentSourceID string `json:"contentSourceId,omitempty"`

	// Content source path
	// Read Only: true
	ContentSourcePath string `json:"contentSourcePath,omitempty"`

	// Content source last sync time
	// Read Only: true
	// Format: date-time
	ContentSourceSyncAt strfmt.DateTime `json:"contentSourceSyncAt,omitempty"`

	// Content source last sync messages
	// Read Only: true
	ContentSourceSyncMessages []string `json:"contentSourceSyncMessages"`

	// Content source last sync status
	// Read Only: true
	// Enum: [SUCCESSFUL FAILED]
	ContentSourceSyncStatus string `json:"contentSourceSyncStatus,omitempty"`

	// Content source type
	// Read Only: true
	ContentSourceType string `json:"contentSourceType,omitempty"`

	// Created time
	// Read Only: true
	// Format: date-time
	CreatedAt strfmt.DateTime `json:"createdAt,omitempty"`

	// Created by
	// Read Only: true
	CreatedBy string `json:"createdBy,omitempty"`

	// Blueprint description
	Description string `json:"description,omitempty"`

	// Object ID
	// Read Only: true
	ID string `json:"id,omitempty"`

	// Blueprint name
	Name string `json:"name,omitempty"`

	// Org ID
	// Read Only: true
	OrgID string `json:"orgId,omitempty"`

	// Project ID
	ProjectID string `json:"projectId,omitempty"`

	// Project Name
	// Read Only: true
	ProjectName string `json:"projectName,omitempty"`

	// Flag to indicate blueprint can be requested from any project in org
	RequestScopeOrg bool `json:"requestScopeOrg,omitempty"`

	// Blueprint self link
	// Read Only: true
	SelfLink string `json:"selfLink,omitempty"`

	// Blueprint status
	// Read Only: true
	// Enum: [DRAFT VERSIONED RELEASED]
	Status string `json:"status,omitempty"`

	// Total released versions
	// Read Only: true
	TotalReleasedVersions int32 `json:"totalReleasedVersions,omitempty"`

	// Total versions
	// Read Only: true
	TotalVersions int32 `json:"totalVersions,omitempty"`

	// Updated time
	// Read Only: true
	// Format: date-time
	UpdatedAt strfmt.DateTime `json:"updatedAt,omitempty"`

	// Updated by
	// Read Only: true
	UpdatedBy string `json:"updatedBy,omitempty"`

	// Validation result on update
	// Read Only: true
	Valid *bool `json:"valid,omitempty"`

	// Validation messages
	// Read Only: true
	ValidationMessages []*BlueprintValidationMessage `json:"validationMessages"`
}


package models 
type SourceControlSyncHistoryItem struct {

	// Content Full Path
	// Read Only: true
	ContentFullPath string `json:"contentFullPath,omitempty"`

	// Content Name
	// Read Only: true
	ContentName string `json:"contentName,omitempty"`

	// Content Type
	// Read Only: true
	// Enum: [BLUEPRINT IMAGE ABX_SCRIPTS TERRAFORM_CONFIGURATION]
	ContentType string `json:"contentType,omitempty"`

	// Details
	// Read Only: true
	Details []string `json:"details"`

	// Unique Id for the Source Control History
	// Format: uuid
	ID strfmt.UUID `json:"id,omitempty"`

	// Integration Id
	// Read Only: true
	IntegrationID string `json:"integrationId,omitempty"`

	// Project Id
	// Read Only: true
	ProjectID string `json:"projectId,omitempty"`

	// Project Name
	// Read Only: true
	ProjectName string `json:"projectName,omitempty"`

	// Id of the sync request
	// Format: uuid
	RequestID strfmt.UUID `json:"requestId,omitempty"`

	// Content source Id
	// Read Only: true
	// Format: uuid
	SourceID strfmt.UUID `json:"sourceId,omitempty"`

	// Status
	// Read Only: true
	Status string `json:"status,omitempty"`

	// Timestamp
	// Read Only: true
	// Format: date-time
	Timestamp strfmt.DateTime `json:"timestamp,omitempty"`
}


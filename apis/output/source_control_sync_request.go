package models 
type SourceControlSyncRequest struct {

	// Created at
	// Read Only: true
	// Format: date-time
	CreatedAt strfmt.DateTime `json:"createdAt,omitempty"`

	// Last Updated at
	// Read Only: true
	// Format: date-time
	LastUpdatedAt strfmt.DateTime `json:"lastUpdatedAt,omitempty"`

	// Message
	// Read Only: true
	Message string `json:"message,omitempty"`

	// Project Id
	// Read Only: true
	ProjectID string `json:"projectId,omitempty"`

	// Request Id
	// Read Only: true
	// Format: uuid
	RequestID strfmt.UUID `json:"requestId,omitempty"`

	// Content Source Id
	// Format: uuid
	SourceID strfmt.UUID `json:"sourceId,omitempty"`

	// Status
	// Read Only: true
	// Enum: [REQUESTED STARTED PROCESSING COMPLETED FAILED SKIPPED]
	Status string `json:"status,omitempty"`
}


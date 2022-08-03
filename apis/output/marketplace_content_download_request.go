package models 
type MarketplaceContentDownloadRequest struct {

	// Content Id
	// Required: true
	ContentID *string `json:"contentId"`

	// Content Name
	// Required: true
	ContentName *string `json:"contentName"`

	// Content Type
	// Required: true
	// Enum: [BLUEPRINT IMAGE ABX_SCRIPTS TERRAFORM_CONFIGURATION]
	ContentType *string `json:"contentType"`

	// Request Created Time
	// Read Only: true
	// Format: date-time
	CreatedAt strfmt.DateTime `json:"createdAt,omitempty"`

	// Downloaded Content Id
	// Read Only: true
	DownloadedContentID string `json:"downloadedContentId,omitempty"`

	// Execution Messages
	// Read Only: true
	ExecutionMessages []*ExecutionMessage `json:"executionMessages"`

	// Request Id
	// Read Only: true
	// Format: uuid
	ID strfmt.UUID `json:"id,omitempty"`

	// Request Last Updated Time
	// Read Only: true
	// Format: date-time
	LastUpdatedAt strfmt.DateTime `json:"lastUpdatedAt,omitempty"`

	// Content Source Id
	// Required: true
	// Format: uuid
	SourceID *strfmt.UUID `json:"sourceId"`

	// Request Status
	// Read Only: true
	Status string `json:"status,omitempty"`

	// Target Id
	// Required: true
	TargetID *string `json:"targetId"`

	// Target Type
	// Required: true
	// Enum: [PROJECT]
	TargetType *string `json:"targetType"`
}


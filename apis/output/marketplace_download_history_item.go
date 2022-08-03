package models 
type MarketplaceDownloadHistoryItem struct {

	// Downloaded Content Id
	ContentID string `json:"contentId,omitempty"`

	// Downloaded Content Name
	ContentName string `json:"contentName,omitempty"`

	// Downloaded Content Type
	// Enum: [BLUEPRINT IMAGE ABX_SCRIPTS TERRAFORM_CONFIGURATION]
	ContentType string `json:"contentType,omitempty"`

	// Downloaded On
	// Format: date-time
	DownloadedOn strfmt.DateTime `json:"downloadedOn,omitempty"`

	// History Id
	ID string `json:"id,omitempty"`

	// Download Status
	// Enum: [FAILED ADDED]
	Status string `json:"status,omitempty"`

	// Target Id
	TargetID string `json:"targetId,omitempty"`

	// Target Name
	TargetName string `json:"targetName,omitempty"`

	// Target Type
	// Enum: [PROJECT]
	TargetType string `json:"targetType,omitempty"`

	// Updated On
	// Format: date-time
	UpdatedOn strfmt.DateTime `json:"updatedOn,omitempty"`

	// Username
	Username string `json:"username,omitempty"`
}


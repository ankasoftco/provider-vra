package models 
type SourceControlSyncRequests struct {

	// content
	Content []*SourceControlSyncRequest `json:"content"`

	// links
	Links []*Link `json:"links"`

	// page
	Page *PageMetadata `json:"page,omitempty"`
}


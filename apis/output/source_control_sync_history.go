package models 
type SourceControlSyncHistory struct {

	// content
	Content []*SourceControlSyncHistoryItem `json:"content"`

	// links
	Links []*Link `json:"links"`

	// page
	Page *PageMetadata `json:"page,omitempty"`
}


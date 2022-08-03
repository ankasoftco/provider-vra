package models 
type ContentSources struct {

	// content
	Content []*ContentSource `json:"content"`

	// links
	Links []*Link `json:"links"`

	// page
	Page *PageMetadata `json:"page,omitempty"`
}


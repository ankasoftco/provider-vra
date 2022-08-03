package models 
type MarketplaceFilterEntries struct {

	// content
	Content []*MarketplaceFilterEntry `json:"content"`

	// links
	Links []*Link `json:"links"`

	// page
	Page *PageMetadata `json:"page,omitempty"`
}


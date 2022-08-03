package models 
type MarketplaceContents struct {

	// content
	Content []*MarketplaceContent `json:"content"`

	// links
	Links []*Link `json:"links"`

	// page
	Page *PageMetadata `json:"page,omitempty"`
}


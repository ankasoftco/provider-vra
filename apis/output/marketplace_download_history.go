package models 
type MarketplaceDownloadHistory struct {

	// content
	Content []*MarketplaceDownloadHistoryItem `json:"content"`

	// links
	Links []*Link `json:"links"`

	// page
	Page *PageMetadata `json:"page,omitempty"`
}


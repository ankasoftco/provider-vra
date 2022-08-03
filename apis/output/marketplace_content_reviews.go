package models 
type MarketplaceContentReviews struct {

	// content
	Content []*MarketplaceContentReview `json:"content"`

	// links
	Links []*Link `json:"links"`

	// page
	Page *PageMetadata `json:"page,omitempty"`
}


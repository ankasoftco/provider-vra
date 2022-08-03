package models 
type MarketplaceContentReview struct {

	// Comment
	Comment string `json:"comment,omitempty"`

	// Content ID
	ContentID string `json:"contentId,omitempty"`

	// Review ID
	ID string `json:"id,omitempty"`

	// Rating
	Rating string `json:"rating,omitempty"`

	// Title
	Title string `json:"title,omitempty"`

	// User
	User string `json:"user,omitempty"`
}


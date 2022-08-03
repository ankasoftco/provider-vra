package models 
type MarketplaceContent struct {

	// Description
	Description string `json:"description,omitempty"`

	// Developer
	Developer string `json:"developer,omitempty"`

	// Display Name
	DisplayName string `json:"displayName,omitempty"`

	// Documents
	Documents []*Document `json:"documents"`

	// EULA
	Eula string `json:"eula,omitempty"`

	// Highlights
	Highlights []string `json:"highlights"`

	// Icon
	Icon string `json:"icon,omitempty"`

	// ID
	ID string `json:"id,omitempty"`

	// Overview
	Overview string `json:"overview,omitempty"`

	// Rating
	Rating string `json:"rating,omitempty"`

	// Screenshots
	Screenshots []string `json:"screenshots"`

	// Short Name
	ShortName string `json:"shortName,omitempty"`

	// support
	Support *Support `json:"support,omitempty"`

	// Tech Specs
	TechSpecs string `json:"techSpecs,omitempty"`

	// Total Reviews
	TotalReviews int32 `json:"totalReviews,omitempty"`

	// Updated Date
	// Format: date-time
	UpdatedAt strfmt.DateTime `json:"updatedAt,omitempty"`

	// Version
	Version string `json:"version,omitempty"`

	// Video
	Video string `json:"video,omitempty"`
}


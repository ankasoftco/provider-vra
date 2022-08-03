package models 
type MarketplaceFilterEntry struct {

	// Entry count for this filter
	// Example: 0
	Count int64 `json:"count,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// Filter Entry Display Name
	Name string `json:"name,omitempty"`
}


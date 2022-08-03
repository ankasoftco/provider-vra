package models 
type PageMetadata struct {

	// number
	Number int64 `json:"number,omitempty"`

	// size
	Size int64 `json:"size,omitempty"`

	// total elements
	TotalElements int64 `json:"totalElements,omitempty"`

	// total pages
	TotalPages int64 `json:"totalPages,omitempty"`
}


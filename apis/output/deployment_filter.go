package models 
type DeploymentFilter struct {

	// content
	Content []*FilterEntry `json:"content"`

	// empty
	Empty bool `json:"empty,omitempty"`

	// first
	First bool `json:"first,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// last
	Last bool `json:"last,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// number
	Number int32 `json:"number,omitempty"`

	// number of elements
	NumberOfElements int32 `json:"numberOfElements,omitempty"`

	// size
	Size int32 `json:"size,omitempty"`

	// sort
	Sort *Sort `json:"sort,omitempty"`

	// total elements
	TotalElements int64 `json:"totalElements,omitempty"`

	// total pages
	TotalPages int32 `json:"totalPages,omitempty"`

	// type
	// Enum: [MULTISELECT DATE_RANGE BOOLEAN]
	Type string `json:"type,omitempty"`
}


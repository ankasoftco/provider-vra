package models 
type SliceOfEventLog struct {

	// content
	Content []*EventLog `json:"content"`

	// empty
	Empty bool `json:"empty,omitempty"`

	// first
	First bool `json:"first,omitempty"`

	// last
	Last bool `json:"last,omitempty"`

	// number
	Number int32 `json:"number,omitempty"`

	// number of elements
	NumberOfElements int32 `json:"numberOfElements,omitempty"`

	// size
	Size int32 `json:"size,omitempty"`

	// sort
	Sort *Sort `json:"sort,omitempty"`
}


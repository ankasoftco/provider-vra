package models 
type StorageSpec struct {

	// desc
	Desc string `json:"desc,omitempty"`

	// document self link
	DocumentSelfLink string `json:"documentSelfLink,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// limit
	Limit int64 `json:"limit,omitempty"`

	// name
	Name string `json:"name,omitempty"`
}


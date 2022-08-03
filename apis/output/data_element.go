package models 
type DataElement struct {

	// description
	Description string `json:"description,omitempty"`

	// error
	Error string `json:"error,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// properties
	Properties interface{} `json:"properties,omitempty"`
}


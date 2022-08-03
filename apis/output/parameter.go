package models 
type Parameter struct {

	// description
	Description string `json:"description,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// param def
	// Enum: [QUERY BODY CONSUMES PRODUCES RESPONSE PATH]
	ParamDef string `json:"paramDef,omitempty"`

	// required
	Required bool `json:"required,omitempty"`

	// type
	Type string `json:"type,omitempty"`

	// value
	Value string `json:"value,omitempty"`
}


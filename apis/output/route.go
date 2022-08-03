package models 
type Route struct {

	// action
	// Enum: [GET POST PATCH PUT DELETE OPTIONS]
	Action string `json:"action,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// handler
	Handler *Operation `json:"handler,omitempty"`

	// matcher
	Matcher *Operation `json:"matcher,omitempty"`

	// parameters
	Parameters []*Parameter `json:"parameters"`

	// path
	Path string `json:"path,omitempty"`

	// support level
	// Enum: [NOT_SUPPORTED INTERNAL DEPRECATED PUBLIC]
	SupportLevel string `json:"supportLevel,omitempty"`
}


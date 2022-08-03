package models 
type NodeList struct {

	// api version
	APIVersion string `json:"apiVersion,omitempty"`

	// items
	Items []*Node `json:"items"`

	// kind
	Kind string `json:"kind,omitempty"`
}


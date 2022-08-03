package models 
type NamedMetering struct {

	// metering
	Metering *Metering `json:"metering,omitempty"`

	// name
	Name string `json:"name,omitempty"`
}


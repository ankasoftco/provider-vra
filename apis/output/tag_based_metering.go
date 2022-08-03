package models 
type TagBasedMetering struct {

	// key
	Key string `json:"key,omitempty"`

	// metering
	Metering *Metering `json:"metering,omitempty"`

	// value
	Value string `json:"value,omitempty"`
}


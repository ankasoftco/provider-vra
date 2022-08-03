package models 
type TagBasedOneTimeMetering struct {

	// key
	Key string `json:"key,omitempty"`

	// one time metering
	OneTimeMetering *OneTimeMetering `json:"oneTimeMetering,omitempty"`

	// value
	Value string `json:"value,omitempty"`
}


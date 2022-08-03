package models 
type Tag struct {

	// Tag's key.
	// Required: true
	Key *string `json:"key"`

	// Tag's value.
	// Required: true
	Value *string `json:"value"`
}


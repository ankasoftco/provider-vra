package models 
type ResourceActionRequest struct {

	// The id of the action to perform.
	ActionID string `json:"actionId,omitempty"`

	// Resource action request inputs
	Inputs interface{} `json:"inputs,omitempty"`

	// Reason for requesting a day2 operation
	Reason string `json:"reason,omitempty"`
}


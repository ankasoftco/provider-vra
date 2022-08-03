package models 
type ResourceResponse struct {

	// failure message
	FailureMessage string `json:"failureMessage,omitempty"`

	// progress message
	ProgressMessage string `json:"progressMessage,omitempty"`

	// resource link
	ResourceLink string `json:"resourceLink,omitempty"`

	// resource properties
	ResourceProperties interface{} `json:"resourceProperties,omitempty"`

	// resource request Id
	ResourceRequestID string `json:"resourceRequestId,omitempty"`

	// status
	// Enum: [CREATED STARTED FINISHED FAILED CANCELLED UNKNOWN]
	Status string `json:"status,omitempty"`
}


package models 
type K8sClusterState struct {

	// error messages
	ErrorMessages string `json:"errorMessages,omitempty"`

	// error reported
	ErrorReported bool `json:"errorReported,omitempty"`

	// error type
	ErrorType string `json:"errorType,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// type
	Type string `json:"type,omitempty"`
}


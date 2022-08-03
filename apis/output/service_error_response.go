package models 
type ServiceErrorResponse struct {

	// details
	Details []string `json:"details"`

	// document kind
	DocumentKind string `json:"documentKind,omitempty"`

	// error code
	ErrorCode int32 `json:"errorCode,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// message Id
	MessageID string `json:"messageId,omitempty"`

	// server error Id
	ServerErrorID string `json:"serverErrorId,omitempty"`

	// stack trace
	StackTrace []string `json:"stackTrace"`

	// status code
	StatusCode int32 `json:"statusCode,omitempty"`
}


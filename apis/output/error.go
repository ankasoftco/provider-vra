package models 
type Error struct {

	// Error message
	// Example: Failed to validate credentials.
	Message string `json:"message,omitempty"`

	// status code
	StatusCode int32 `json:"statusCode,omitempty"`
}


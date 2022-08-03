package models 
type AuthResponse struct {

	// Base64 encoded auth token.
	// Required: true
	Token *string `json:"token"`

	// Type of the token.
	// Example: Bearer
	// Required: true
	TokenType *string `json:"tokenType"`
}


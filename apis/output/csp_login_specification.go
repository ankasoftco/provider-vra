package models 
type CspLoginSpecification struct {

	// Refresh token obtained from the UI
	// Example: 5e7c2c-9a9e-4b0-9339-a7f94
	// Required: true
	RefreshToken *string `json:"refreshToken"`
}


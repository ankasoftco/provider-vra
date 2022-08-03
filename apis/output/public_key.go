package models 
type PublicKey struct {

	// algorithm
	Algorithm string `json:"algorithm,omitempty"`

	// encoded
	// Format: byte
	Encoded strfmt.Base64 `json:"encoded,omitempty"`

	// format
	Format string `json:"format,omitempty"`
}


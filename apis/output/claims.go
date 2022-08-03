package models 
type Claims struct {

	// audience
	Audience []string `json:"audience"`

	// expiration time
	ExpirationTime int64 `json:"expirationTime,omitempty"`

	// issued at
	IssuedAt int64 `json:"issuedAt,omitempty"`

	// issuer
	Issuer string `json:"issuer,omitempty"`

	// jwt Id
	JwtID string `json:"jwtId,omitempty"`

	// not before
	NotBefore int64 `json:"notBefore,omitempty"`

	// properties
	Properties map[string]string `json:"properties,omitempty"`

	// subject
	Subject string `json:"subject,omitempty"`
}


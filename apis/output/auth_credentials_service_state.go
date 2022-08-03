package models 
type AuthCredentialsServiceState struct {

	// bearer
	Bearer bool `json:"bearer,omitempty"`

	// custom properties
	CustomProperties map[string]string `json:"customProperties,omitempty"`

	// document self link
	DocumentSelfLink string `json:"documentSelfLink,omitempty"`

	// password
	Password bool `json:"password,omitempty"`

	// private key
	PrivateKey string `json:"privateKey,omitempty"`

	// private key Id
	PrivateKeyID string `json:"privateKeyId,omitempty"`

	// public key
	PublicKey string `json:"publicKey,omitempty"`

	// tenant links
	TenantLinks []string `json:"tenantLinks"`

	// type
	Type string `json:"type,omitempty"`

	// user email
	UserEmail string `json:"userEmail,omitempty"`
}


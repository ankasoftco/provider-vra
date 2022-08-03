package models 
type VcfServiceAccount struct {

	// auth credentials
	AuthCredentials string `json:"authCredentials,omitempty"`

	// credential Id
	CredentialID string `json:"credentialId,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// password
	Password string `json:"password,omitempty"`

	// resource Id
	ResourceID string `json:"resourceId,omitempty"`

	// resource type
	ResourceType string `json:"resourceType,omitempty"`

	// role name
	RoleName string `json:"roleName,omitempty"`

	// username
	Username string `json:"username,omitempty"`
}


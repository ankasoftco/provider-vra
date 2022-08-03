package models 
type TerraformVersion struct {

	// The type of authentication for the download url
	// Enum: [NONE BASIC]
	AuthenticationType string `json:"authenticationType,omitempty"`

	// Created time
	// Read Only: true
	// Format: date-time
	CreatedAt strfmt.DateTime `json:"createdAt,omitempty"`

	// Created by
	// Read Only: true
	CreatedBy string `json:"createdBy,omitempty"`

	// Version description
	Description string `json:"description,omitempty"`

	// Version status
	Enabled bool `json:"enabled,omitempty"`

	// Version ID
	// Read Only: true
	// Format: uuid
	ID strfmt.UUID `json:"id,omitempty"`

	// Org ID
	// Read Only: true
	OrgID string `json:"orgId,omitempty"`

	// The password for basic authentication
	Password string `json:"password,omitempty"`

	// The sha256 checksum of the terraform binary
	Sha256Checksum string `json:"sha256Checksum,omitempty"`

	// Updated time
	// Read Only: true
	// Format: date-time
	UpdatedAt strfmt.DateTime `json:"updatedAt,omitempty"`

	// Updated by
	// Read Only: true
	UpdatedBy string `json:"updatedBy,omitempty"`

	// Download url
	URL string `json:"url,omitempty"`

	// The user name for basic authentication
	Username string `json:"username,omitempty"`

	// The numeric version of terraform release
	Version string `json:"version,omitempty"`
}


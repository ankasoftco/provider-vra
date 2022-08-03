package models 
type AboutDeprecationPolicy struct {

	// The date the api was deprecated in yyyy-MM-dd format (UTC). Could be empty if the api is not deprecated.
	DeprecatedAt string `json:"deprecatedAt,omitempty"`

	// A free text description that contains information about why this api is deprecated and how to migrate to a newer version.
	Description string `json:"description,omitempty"`

	// The date the api support will be dropped in yyyy-MM-dd format (UTC). The api may still be available for use after that date but this is not guaranteed.
	ExpiresAt string `json:"expiresAt,omitempty"`
}


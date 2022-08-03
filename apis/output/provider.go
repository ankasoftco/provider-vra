package models 
type Provider struct {

	// The Terraform provider's statically configured attributes.
	Attributes map[string]string `json:"attributes,omitempty"`

	// The Terraform provider's type.
	Type string `json:"type,omitempty"`
}


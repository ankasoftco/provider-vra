package models 
type OutputValue struct {

	// The description of the output value in the Terraform configuration.
	Description string `json:"description,omitempty"`

	// The name of the output value in the Terraform configuration.
	Name string `json:"name,omitempty"`

	// A flag indicating that the field should be obscured because of security concerns.
	Sensitive bool `json:"sensitive,omitempty"`
}


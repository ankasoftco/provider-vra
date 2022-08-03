package models 
type Variable struct {

	// The variable's default value in the Terraform configuration.
	DefaultValue interface{} `json:"defaultValue,omitempty"`

	// The variable's description in the Terraform configuration.
	Description string `json:"description,omitempty"`

	// The variable's name in the Terraform configuration.
	Name string `json:"name,omitempty"`

	// Whether the variable should be obscured because of security concerns.
	Sensitive bool `json:"sensitive,omitempty"`

	// The variable's type in the Terraform configuration. Complex Terraform types may be treated as Strings.
	// Enum: [STRING NUMBER BOOL LIST MAP]
	Type string `json:"type,omitempty"`
}


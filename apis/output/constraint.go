package models 
type Constraint struct {

	// An expression of the form "[!]tag-key[:[tag-value]]", used to indicate a constraint match on keys and values of tags.
	//
	// Example: ha:strong
	// Required: true
	Expression *string `json:"expression"`

	// Indicates whether this constraint should be strictly enforced or not.
	// Required: true
	Mandatory *bool `json:"mandatory"`
}


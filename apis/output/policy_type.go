package models 
type PolicyType struct {

	// Defines configuration options for policy type
	Config *PolicyFeatureConfig `json:"config,omitempty"`

	// Schema describing a policy object of this type
	// Required: true
	DefinitionSchema interface{} `json:"definitionSchema"`

	// Policy type display-name/label
	// Required: true
	DisplayName *string `json:"displayName"`

	// Policy type ID
	// Required: true
	ID *string `json:"id"`

	// Policy type name
	// Required: true
	Name *string `json:"name"`

	// Schema describing objects that can be affected by this policy
	// Required: true
	TargetSchema interface{} `json:"targetSchema"`
}


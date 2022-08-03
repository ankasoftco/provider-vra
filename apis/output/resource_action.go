package models 
type ResourceAction struct {

	// Resource action type
	// Enum: [RESOURCE_ACTION RESOURCE_EXTENSION]
	ActionType string `json:"actionType,omitempty"`

	// Dependent resources
	Dependents []string `json:"dependents"`

	// Resource action description
	Description string `json:"description,omitempty"`

	// Resource action display name
	DisplayName string `json:"displayName,omitempty"`

	// Resource action custom UI definition. Optional
	FormDefinition *FormDefinition `json:"formDefinition,omitempty"`

	// Resource action id
	ID string `json:"id,omitempty"`

	// Resource action name
	Name string `json:"name,omitempty"`

	// Resource action org ID
	OrgID string `json:"orgId,omitempty"`

	// Resource action project ID
	ProjectID string `json:"projectId,omitempty"`

	// Resource action input schema
	Schema interface{} `json:"schema,omitempty"`

	// Resource action is valid for current state
	Valid bool `json:"valid,omitempty"`
}


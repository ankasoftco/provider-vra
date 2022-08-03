package models 
type BlueprintContentSource struct {

	// config
	Config *BlueprintContentSourceConfig `json:"config,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// integration account name
	IntegrationAccountName string `json:"integrationAccountName,omitempty"`

	// is sync enabled
	IsSyncEnabled bool `json:"isSyncEnabled,omitempty"`

	// label
	Label string `json:"label,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// project Id
	ProjectID string `json:"projectId,omitempty"`

	// type Id
	TypeID string `json:"typeId,omitempty"`
}


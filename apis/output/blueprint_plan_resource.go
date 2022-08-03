package models 
type BlueprintPlanResource struct {

	// Resource depends on other resources in the plan
	// Read Only: true
	DependsOnResources []string `json:"dependsOnResources"`

	// Resource new properties
	// Read Only: true
	NewProperties interface{} `json:"newProperties,omitempty"`

	// Resource old properties
	// Read Only: true
	OldProperties interface{} `json:"oldProperties,omitempty"`

	// Resource name
	// Read Only: true
	ResourceName string `json:"resourceName,omitempty"`

	// Resource reason
	// Read Only: true
	// Enum: [CREATE RECREATE UPDATE DELETE ACTION READ]
	ResourceReason string `json:"resourceReason,omitempty"`

	// Resource type
	// Read Only: true
	ResourceType string `json:"resourceType,omitempty"`

	// List of task names
	// Read Only: true
	TaskNames []string `json:"taskNames"`
}


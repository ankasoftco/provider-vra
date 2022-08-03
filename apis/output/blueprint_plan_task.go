package models 
type BlueprintPlanTask struct {

	// Tasks depends on other tasks
	// Read Only: true
	DependsOnTasks []string `json:"dependsOnTasks"`

	// Task input properties
	// Read Only: true
	InputProperties interface{} `json:"inputProperties,omitempty"`

	// Task name
	// Read Only: true
	Name string `json:"name,omitempty"`

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
}


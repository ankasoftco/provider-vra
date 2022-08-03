package models 
type BlueprintPlan struct {

	// Blueprint plan input properties
	// Read Only: true
	InputProperties interface{} `json:"inputProperties,omitempty"`

	// Blueprint plan output properties
	// Read Only: true
	OutputProperties interface{} `json:"outputProperties,omitempty"`

	// Blueprint plan tasks
	// Read Only: true
	Tasks []*BlueprintPlanTask `json:"tasks"`
}


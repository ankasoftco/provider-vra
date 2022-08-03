package models 
type BlueprintResourcesPlan struct {

	// Blueprint plan tasks
	// Read Only: true
	Resources []*BlueprintPlanResource `json:"resources"`
}


package models 
type MeteringPolicyAssignmentInfo struct {

	// count
	Count int32 `json:"count,omitempty"`

	// entity type
	// Enum: [ALL PROJECT CLOUDZONE]
	EntityType string `json:"entityType,omitempty"`
}


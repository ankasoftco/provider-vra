package models 
type SupervisorNamespaceSyncResponseDTO struct {

	// message
	Message string `json:"message,omitempty"`

	// operation tracker link
	OperationTrackerLink string `json:"operationTrackerLink,omitempty"`

	// status
	// Enum: [IN_PROGRESS COMPLETED FAILED]
	Status string `json:"status,omitempty"`
}


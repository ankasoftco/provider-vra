package models 
type TileExecutorResponse struct {

	// time taken for the Tile execution to finish.
	Duration int64 `json:"duration,omitempty"`

	// Output properties from the Tile execution.
	Output interface{} `json:"output,omitempty"`

	// Status of the Tile execution.
	// Example: COMPLETED
	// Enum: [CREATED STARTED PAUSED CANCELED COMPLETED FAILED]
	Status string `json:"status,omitempty"`

	// Status message of the Tile execution.
	// Example: Failed to connect to jenkins server endpoint
	StatusMessage string `json:"statusMessage,omitempty"`
}


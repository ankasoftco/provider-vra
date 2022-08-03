package models 
type GerritTriggerGerritEventConfiguration struct {

	// Type of the gerrit event.
	// Example: patchset-created
	EventType string `json:"eventType,omitempty"`

	// Comment to be posted to the ChangeSet on execution termination.
	// Example: Pipeline Execution Failed
	FailureComment string `json:"failureComment,omitempty"`

	// Map representing the Input properties for the Pipeline.
	// Example: [{"ip":"10.5.23.84","script":"testScript.sh"}]
	Input map[string]string `json:"input,omitempty"`

	// Pipeline that needs to be triggered on receiving this event.
	// Example: DemoPipeline
	Pipeline string `json:"pipeline,omitempty"`

	// Comment to be posted to the ChangeSet on execution termination.
	// Example: Pipeline Execution Completed
	SuccessComment string `json:"successComment,omitempty"`

	// The label to be posted on Gerrit Server to perform actions.
	// Example: Verified +1
	VerifiedLabel string `json:"verifiedLabel,omitempty"`
}


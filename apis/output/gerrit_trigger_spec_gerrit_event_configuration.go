package models 
type GerritTriggerSpecGerritEventConfiguration struct {

	// Type of the gerrit event.
	// Example: patchset-created
	// Required: true
	EventType *string `json:"eventType"`

	// Comment to be posted to the ChangeSet on execution termination.
	// Example: Pipeline Execution Failed
	FailureComment string `json:"failureComment,omitempty"`

	// Map representing the Input properties for the Pipeline.
	// Example: [{"ip":"10.5.23.84","script":"testScript.sh"}]
	Input map[string]string `json:"input,omitempty"`

	// Pipeline that needs to be triggered on receiving this event.
	// Example: DemoPipeline
	// Required: true
	Pipeline *string `json:"pipeline"`

	// Comment to be posted to the ChangeSet on execution termination.
	// Example: Pipeline Execution Completed
	SuccessComment string `json:"successComment,omitempty"`

	// The label to be posted on Gerrit Server to perform actions.
	// Example: Verified +1
	VerifiedLabel string `json:"verifiedLabel,omitempty"`
}


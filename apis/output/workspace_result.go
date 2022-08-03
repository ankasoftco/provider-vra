package models 
type WorkspaceResult struct {

	// Logs for the executing step.
	Logs []string `json:"logs"`

	// Status of the Git clone/creating a container for the Workspace.
	// Example: COMPLETED
	Status string `json:"status,omitempty"`

	// The current step for Git clone/creating a container for the Workspace.
	// Example: GIT_CLONE
	Step string `json:"step,omitempty"`
}


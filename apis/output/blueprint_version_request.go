package models 
type BlueprintVersionRequest struct {

	// Blueprint version change log
	ChangeLog string `json:"changeLog,omitempty"`

	// Blueprint version description
	Description string `json:"description,omitempty"`

	// Flag indicating to release version
	Release bool `json:"release,omitempty"`

	// Blueprint version
	// Required: true
	Version *string `json:"version"`
}


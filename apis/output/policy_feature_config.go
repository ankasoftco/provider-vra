package models 
type PolicyFeatureConfig struct {

	// enable dry run
	EnableDryRun bool `json:"enableDryRun,omitempty"`

	// enable enforcement type
	EnableEnforcementType bool `json:"enableEnforcementType,omitempty"`

	// enable update notification
	EnableUpdateNotification bool `json:"enableUpdateNotification,omitempty"`
}


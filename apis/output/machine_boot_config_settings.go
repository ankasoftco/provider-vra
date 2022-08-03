package models 
type MachineBootConfigSettings struct {

	// In case a timeout occurs whether the provisioning process should fail or continue.
	// Example: false
	PhoneHomeFailOnTimeout bool `json:"phoneHomeFailOnTimeout,omitempty"`

	// A phone_home module will be added to the Cloud Config and the provisioning will wait on a callback prior proceeding
	// Example: true
	PhoneHomeShouldWait bool `json:"phoneHomeShouldWait,omitempty"`

	// The period of time to wait for the phone_home module callback to occur
	// Example: 100
	PhoneHomeTimeoutSeconds int32 `json:"phoneHomeTimeoutSeconds,omitempty"`
}


package models 
type CustomNamingTemplate struct {

	// counters
	Counters []*CustomNamingCounter `json:"counters"`

	// Unique id of custom naming template
	// Example: 3fa85f64-5717-4562-b3fc-2c963f76afa6
	// Format: uuid
	ID strfmt.UUID `json:"id,omitempty"`

	// Set the increment to the counter value to be taken for each name.
	// Example: 1
	IncrementStep int64 `json:"incrementStep,omitempty"`

	// The specified template used to generate the resource names
	// Example: mcm-${project.name}-${##}
	Pattern string `json:"pattern,omitempty"`

	// Flag to represent default pattern or static pattern
	// Example: true
	ResourceDefault bool `json:"resourceDefault,omitempty"`

	// Resource type
	// Enum: [COMPUTE NETWORK COMPUTE_STORAGE LOAD_BALANCER RESOURCE_GROUP GATEWAY NAT SECURITY_GROUP]
	ResourceType string `json:"resourceType,omitempty"`

	// Resource type
	// Example: Network
	ResourceTypeName string `json:"resourceTypeName,omitempty"`

	// The value from which naming pattern counter will start.
	// Example: 2
	StartCounter int64 `json:"startCounter,omitempty"`

	// Static pattern text
	// Example: mcm-project1-
	StaticPattern string `json:"staticPattern,omitempty"`

	// Flag to check if name should be unique
	// Example: true
	UniqueName bool `json:"uniqueName,omitempty"`
}


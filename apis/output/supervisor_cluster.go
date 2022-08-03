package models 
type SupervisorCluster struct {

	// address
	Address string `json:"address,omitempty"`

	// cpu capacity
	CPUCapacity int64 `json:"cpuCapacity,omitempty"`

	// cpu used
	CPUUsed int64 `json:"cpuUsed,omitempty"`

	// document self link
	DocumentSelfLink string `json:"documentSelfLink,omitempty"`

	// endpoint link
	EndpointLink string `json:"endpointLink,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// memory capacity
	MemoryCapacity int64 `json:"memoryCapacity,omitempty"`

	// memory used
	MemoryUsed int64 `json:"memoryUsed,omitempty"`

	// moref
	Moref string `json:"moref,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// registered
	Registered bool `json:"registered,omitempty"`

	// status
	// Enum: [ON CONFIGURING REMOVING ERROR]
	Status string `json:"status,omitempty"`

	// status message
	StatusMessage string `json:"statusMessage,omitempty"`

	// storage capacity
	StorageCapacity int64 `json:"storageCapacity,omitempty"`

	// storage used
	StorageUsed int64 `json:"storageUsed,omitempty"`

	// type
	Type string `json:"type,omitempty"`

	// updated millis
	UpdatedMillis int64 `json:"updatedMillis,omitempty"`
}


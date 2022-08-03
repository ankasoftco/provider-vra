package models 
type FabricFlavorDescription struct {

	// Number of CPU cores. Mandatory for private clouds such as vSphere. Not populated when inapplicable.
	// Example: 4
	CPUCount int32 `json:"cpuCount,omitempty"`

	// The id of the instance type in the corresponding cloud.
	// Example: 901004, 901008
	ID string `json:"id,omitempty"`

	// Total amount of memory (in megabytes). Mandatory for private clouds such as vSphere. Not populated when inapplicable.
	// Example: 4096
	MemoryInMB int64 `json:"memoryInMB,omitempty"`

	// The value of the instance type in the corresponding cloud. Valid and mandatory for public clouds
	// Example: t2.small, t2.medium
	Name string `json:"name,omitempty"`
}


package models 
type K8SNode struct {

	// allocatable pods
	AllocatablePods int32 `json:"allocatablePods,omitempty"`

	// cpu cores
	CPUCores float64 `json:"cpuCores,omitempty"`

	// cpu usage
	CPUUsage float64 `json:"cpuUsage,omitempty"`

	// created millis
	CreatedMillis int64 `json:"createdMillis,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// id
	// Format: uuid
	ID strfmt.UUID `json:"id,omitempty"`

	// memory usage
	MemoryUsage float64 `json:"memoryUsage,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// org Id
	OrgID string `json:"orgId,omitempty"`

	// total memory
	TotalMemory float64 `json:"totalMemory,omitempty"`

	// updated millis
	UpdatedMillis int64 `json:"updatedMillis,omitempty"`
}


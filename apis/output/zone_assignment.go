package models 
type ZoneAssignment struct {

	// The amount of CPUs currently allocated.
	AllocatedCPU int64 `json:"allocatedCpu,omitempty"`

	// The number of resource instances currently allocated
	AllocatedInstancesCount int64 `json:"allocatedInstancesCount,omitempty"`

	// The amount of memory currently allocated.
	AllocatedMemoryMB int64 `json:"allocatedMemoryMB,omitempty"`

	// The amount of storage currently allocated.
	AllocatedStorageGB float64 `json:"allocatedStorageGB,omitempty"`

	// The maximum amount of cpus that can be used by this cloud zone. Default is 0 (unlimited cpu).
	// Example: 2048
	CPULimit int64 `json:"cpuLimit,omitempty"`

	// The maximum number of instances that can be provisioned in this cloud zone. Default is 0 (unlimited instances).
	// Example: 50
	MaxNumberInstances int64 `json:"maxNumberInstances,omitempty"`

	// The maximum amount of memory that can be used by this cloud zone. Default is 0 (unlimited memory).
	// Example: 2048
	MemoryLimitMB int64 `json:"memoryLimitMB,omitempty"`

	// The priority of this zone in the current project. Lower numbers mean higher priority. Default is 0 (highest)
	// Example: 1
	Priority int32 `json:"priority,omitempty"`

	// Defines an upper limit on storage that can be requested from a cloud zone which is part of this project. Default is 0 (unlimited storage). Please note that this feature is supported only for vSphere cloud zones. Not valid for other cloud zone types.
	// Example: 20
	StorageLimitGB int64 `json:"storageLimitGB,omitempty"`

	// The Cloud Zone Id
	// Example: 77ee1
	ZoneID string `json:"zoneId,omitempty"`
}


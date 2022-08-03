package models 
type FabricFlavor struct {

	// Size of the boot disk (in megabytes). Not populated when inapplicable.
	// Example: 486400
	BootDiskSizeInMB int32 `json:"bootDiskSizeInMB,omitempty"`

	// Number of CPU cores. Not populated when inapplicable.
	// Example: 2
	CPUCount int32 `json:"cpuCount,omitempty"`

	// Number of data disks. Not populated when inapplicable.
	// Example: 1
	DataDiskMaxCount int32 `json:"dataDiskMaxCount,omitempty"`

	// Size of the data disks (in megabytes). Not populated when inapplicable.
	// Example: 486400
	DataDiskSizeInMB int32 `json:"dataDiskSizeInMB,omitempty"`

	// The internal identification used by the corresponding cloud end-point
	// Example: 901004
	ID string `json:"id,omitempty"`

	// Total amount of memory (in megabytes). Not populated when inapplicable.
	// Example: 15616
	MemoryInMB int64 `json:"memoryInMB,omitempty"`

	// The value of the instance type in the corresponding cloud.
	// Example: i3.large
	// Required: true
	Name *string `json:"name"`

	// The type of network supported by this instance type. Not populated when inapplicable.
	// Example: Up to 10 Gigabit
	NetworkType string `json:"networkType,omitempty"`

	// The type of storage supported by this instance type. Not populated when inapplicable.
	// Example: NVMe_SSD
	StorageType string `json:"storageType,omitempty"`
}


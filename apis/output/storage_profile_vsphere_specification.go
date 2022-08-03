package models 
type StorageProfileVsphereSpecification struct {

	// Id of the vSphere Datastore for placing disk and VM.
	// Example: 08d28
	DatastoreID string `json:"datastoreId,omitempty"`

	// Indicates if a storage profile acts as a default storage profile for a disk.
	// Example: true
	// Required: true
	DefaultItem *bool `json:"defaultItem"`

	// A human-friendly description.
	Description string `json:"description,omitempty"`

	// Type of mode for the disk
	// Example: undefined / independent-persistent / independent-nonpersistent
	DiskMode string `json:"diskMode,omitempty"`

	// Disk types are specified as
	//
	//  	Standard - Simple vSphere virtual disks which cannot be managed independently without an attached VM.
	// 	First Class - Improved version of standard virtual disks, designed to be fully mananged independent storage objects.
	//
	// Empty value is considered as Standard
	// Example: standard / firstClass
	DiskType string `json:"diskType,omitempty"`

	// The upper bound for the I/O operations per second allocated for each virtual disk.
	// Example: 1000
	LimitIops string `json:"limitIops,omitempty"`

	// A human-friendly name used as an identifier in APIs that support this option.
	// Required: true
	Name *string `json:"name"`

	// Type of provisioning policy for the disk.
	// Example: thin / thick / eagerZeroedThick
	ProvisioningType string `json:"provisioningType,omitempty"`

	// The Id of the region that is associated with the storage profile.
	// Example: 31186
	// Required: true
	RegionID *string `json:"regionId"`

	// A specific number of shares assigned to each virtual machine.
	// Example: 2000
	Shares string `json:"shares,omitempty"`

	// Shares are specified as High, Normal, Low or Custom and these values specify share values with a 4:2:1 ratio, respectively.
	// Example: low / normal / high / custom
	SharesLevel string `json:"sharesLevel,omitempty"`

	// Id of the vSphere Storage Policy to be applied.
	// Example: 6b59743af31d
	StoragePolicyID string `json:"storagePolicyId,omitempty"`

	// Indicates whether this storage profile supports encryption or not.
	// Example: false
	SupportsEncryption bool `json:"supportsEncryption,omitempty"`

	// A list of tags that represent the capabilities of this storage profile.
	// Example: [ { \"key\" : \"tier\", \"value\": \"silver\" } ]
	Tags []*Tag `json:"tags"`
}


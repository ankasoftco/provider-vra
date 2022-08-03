package models 
type StorageProfileAzureSpecification struct {

	// Indicates the caching mechanism for additional disk.
	// Example: None / ReadOnly / ReadWrite
	DataDiskCaching string `json:"dataDiskCaching,omitempty"`

	// Indicates if a storage policy contains default storage properties.
	// Example: true
	DefaultItem bool `json:"defaultItem,omitempty"`

	// A human-friendly description.
	Description string `json:"description,omitempty"`

	// Indicates the id of disk encryption set.
	// Example: /subscriptions/b8ef63/resourceGroups/DiskEncryptionSets/providers/Microsoft.Compute/diskEncryptionSets/MyDES
	DiskEncryptionSetID string `json:"diskEncryptionSetId,omitempty"`

	// Indicates the performance tier for the storage type. Premium disks are SSD backed and Standard disks are HDD backed.
	// Example: Standard_LRS / Premium_LRS
	DiskType string `json:"diskType,omitempty"`

	// A human-friendly name used as an identifier in APIs that support this option.
	// Required: true
	Name *string `json:"name"`

	// Indicates the caching mechanism for OS disk. Default policy for OS disks is Read/Write.
	// Example: None / ReadOnly / ReadWrite
	OsDiskCaching string `json:"osDiskCaching,omitempty"`

	// The If of the region that is associated with the storage profile.
	// Example: 31186
	// Required: true
	RegionID *string `json:"regionId"`

	// Id of a storage account where in the disk is placed.
	// Example: aaa82
	StorageAccountID string `json:"storageAccountId,omitempty"`

	// Indicates whether this storage policy should support encryption or not.
	// Example: false
	SupportsEncryption bool `json:"supportsEncryption,omitempty"`

	// A set of tag keys and optional values for a storage policy which define set of specifications for creating a disk.
	// Example: [ { \"key\" : \"tier\", \"value\": \"silver\" } ]
	Tags []*Tag `json:"tags"`
}


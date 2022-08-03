package models 
type StorageProfileSpecification struct {

	// Indicates if a storage profile is a default profile.
	// Required: true
	DefaultItem *bool `json:"defaultItem"`

	// A human-friendly description.
	Description string `json:"description,omitempty"`

	// Map of storage properties that are to be applied on disk while provisioning.
	// Example: { \"diskProperties\": {\n                    \"provisioningType\": \"thin\",\n                    \"sharesLevel\": \"low\",\n                    \"shares\": \"500\",\n                    \"limitIops\": \"500\"\n                    \"diskType\": \"firstClass\"\n                    \"deviceType\": \"ebs\"\n                    \"volumeType\": \"gp2\"\n                    \"azureDataDiskCaching\": \"ReadWrite\"\n                    \"azureOsDiskCaching\": \"ReadWrite\"\n                    \"azureManagedDiskType\": \"Standard_LRS\"\n                } }
	DiskProperties map[string]string `json:"diskProperties,omitempty"`

	// Map of storage placements to know where the disk is provisioned.
	// Example: { \"diskTargetProperties\": {\n                    \"storageAccountId\": \"27dhbf7\",\n                    \"storagePolicyId\": \"7fhfj9f\",\n                    \"datastoreId\": \"638nfjd8\",\n                } }
	DiskTargetProperties map[string]string `json:"diskTargetProperties,omitempty"`

	// A human-friendly name used as an identifier in APIs that support this option.
	// Required: true
	Name *string `json:"name"`

	// The Id of the region that is associated with the storage profile.
	// Example: 31186
	// Required: true
	RegionID *string `json:"regionId"`

	// Indicates whether this storage profile supports encryption or not.
	SupportsEncryption bool `json:"supportsEncryption,omitempty"`

	// A list of tags that represent the capabilities of this storage profile
	// Example: [ { \"key\" : \"tier\", \"value\": \"silver\" } ]
	Tags []*Tag `json:"tags"`
}


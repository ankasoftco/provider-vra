package models 
type AzureStorageProfile struct {

	// HATEOAS of the entity
	// Required: true
	Links map[string]Href `json:"_links"`

	// Id of the cloud account this storage profile belongs to.
	// Example: [9e49]
	CloudAccountID string `json:"cloudAccountId,omitempty"`

	// Date when the entity was created. The date is in ISO 8601 and UTC.
	// Example: 2012-09-27
	CreatedAt string `json:"createdAt,omitempty"`

	// Indicates the caching mechanism for additional disk.
	// Example: None / ReadOnly / ReadWrite
	DataDiskCaching string `json:"dataDiskCaching,omitempty"`

	// Indicates if a storage profile contains default storage properties.
	// Example: false
	// Required: true
	DefaultItem *bool `json:"defaultItem"`

	// A human-friendly description.
	// Example: my-description
	Description string `json:"description,omitempty"`

	// Indicates the id of disk encryption set.
	// Example: /subscriptions/b8ef63/resourceGroups/DiskEncryptionSets/providers/Microsoft.Compute/diskEncryptionSets/MyDES
	DiskEncryptionSetID string `json:"diskEncryptionSetId,omitempty"`

	// Indicates the performance tier for the storage type. Premium disks are SSD backed and Standard disks are HDD backed.
	// Example: Standard_LRS / Premium_LRS
	DiskType string `json:"diskType,omitempty"`

	// The id of the region for which this profile is defined
	// Example: uswest
	ExternalRegionID string `json:"externalRegionId,omitempty"`

	// The id of this resource instance
	// Example: 9e49
	// Required: true
	ID *string `json:"id"`

	// A human-friendly name used as an identifier in APIs that support this option.
	// Example: my-name
	Name string `json:"name,omitempty"`

	// The id of the organization this entity belongs to.
	// Example: 42413b31-1716-477e-9a88-9dc1c3cb1cdf
	OrgID string `json:"orgId,omitempty"`

	// Indicates the caching mechanism for OS disk. Default policy for OS disks is Read/Write.
	// Example: None / ReadOnly / ReadWrite
	OsDiskCaching string `json:"osDiskCaching,omitempty"`

	// Email of the user that owns the entity.
	// Example: csp@vmware.com
	Owner string `json:"owner,omitempty"`

	// Indicates whether this storage profile should support encryption or not.
	// Example: false
	SupportsEncryption bool `json:"supportsEncryption,omitempty"`

	// A list of tags that represent the capabilities of this storage profile
	// Example: [ { \"key\" : \"tier\", \"value\": \"silver\" } ]
	Tags []*Tag `json:"tags"`

	// Date when the entity was last updated. The date is ISO 8601 and UTC.
	// Example: 2012-09-27
	UpdatedAt string `json:"updatedAt,omitempty"`
}


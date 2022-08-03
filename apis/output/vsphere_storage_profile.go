package models 
type VsphereStorageProfile struct {

	// HATEOAS of the entity
	// Required: true
	Links map[string]Href `json:"_links"`

	// Id of the cloud account this storage profile belongs to.
	// Example: [9e49]
	CloudAccountID string `json:"cloudAccountId,omitempty"`

	// Date when the entity was created. The date is in ISO 8601 and UTC.
	// Example: 2012-09-27
	CreatedAt string `json:"createdAt,omitempty"`

	// Indicates if a storage profile contains default storage properties.
	// Example: false
	// Required: true
	DefaultItem *bool `json:"defaultItem"`

	// A human-friendly description.
	// Example: my-description
	Description string `json:"description,omitempty"`

	// Type of mode for the disk
	// Example: undefined / independent-persistent / independent-nonpersistent
	DiskMode string `json:"diskMode,omitempty"`

	// Disk types are specified as
	// 	Standard - Simple vSphere virtual disks which cannot be managed independently without an attached VM.
	// 	First Class - Improved version of standard virtual disks, designed to be fully mananged
	//  independent storage objects.
	// Empty value is considered as Standard
	// Example: firstClass / standard
	DiskType string `json:"diskType,omitempty"`

	// The id of the region for which this profile is defined
	// Example: Datacenter:datacenter-2
	ExternalRegionID string `json:"externalRegionId,omitempty"`

	// The id of this resource instance
	// Example: 9e49
	// Required: true
	ID *string `json:"id"`

	// The upper bound for the I/O operations per second allocated for each disk.
	// Example: 1000
	LimitIops string `json:"limitIops,omitempty"`

	// A human-friendly name used as an identifier in APIs that support this option.
	// Example: my-name
	Name string `json:"name,omitempty"`

	// The id of the organization this entity belongs to.
	// Example: 42413b31-1716-477e-9a88-9dc1c3cb1cdf
	OrgID string `json:"orgId,omitempty"`

	// Email of the user that owns the entity.
	// Example: csp@vmware.com
	Owner string `json:"owner,omitempty"`

	// Type of format for the disk.
	// Example: thin / thick / eagerZeroedThick
	ProvisioningType string `json:"provisioningType,omitempty"`

	// A specific number of shares assigned to each virtual machine.
	// Example: 2000
	Shares string `json:"shares,omitempty"`

	// Shares level are specified as High, Normal, Low or Custom.
	// Example: low / normal / high / custom
	SharesLevel string `json:"sharesLevel,omitempty"`

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


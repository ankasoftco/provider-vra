package models 
type AwsStorageProfile struct {

	// HATEOAS of the entity
	// Required: true
	Links map[string]Href `json:"_links"`

	// Id of the cloud account this storage profile belongs to.
	// Example: [9e49]
	CloudAccountID string `json:"cloudAccountId,omitempty"`

	// Date when the entity was created. The date is in ISO 8601 and UTC.
	// Example: 2012-09-27
	CreatedAt string `json:"createdAt,omitempty"`

	// Indicates whether this storage profile is default or not..
	// Example: false
	// Required: true
	DefaultItem *bool `json:"defaultItem"`

	// A human-friendly description.
	// Example: my-description
	Description string `json:"description,omitempty"`

	// Indicates the type of storage device.
	// Example: ebs / instance-store
	DeviceType string `json:"deviceType,omitempty"`

	// The id of the region for which this profile is defined
	// Example: us-east-1
	ExternalRegionID string `json:"externalRegionId,omitempty"`

	// The id of this resource instance
	// Example: 9e49
	// Required: true
	ID *string `json:"id"`

	// Indicates maximum I/O operations per second in range(1-20,000).
	// Example: 2000
	Iops string `json:"iops,omitempty"`

	// A human-friendly name used as an identifier in APIs that support this option.
	// Example: my-name
	Name string `json:"name,omitempty"`

	// The id of the organization this entity belongs to.
	// Example: 42413b31-1716-477e-9a88-9dc1c3cb1cdf
	OrgID string `json:"orgId,omitempty"`

	// Email of the user that owns the entity.
	// Example: csp@vmware.com
	Owner string `json:"owner,omitempty"`

	// Indicates whether this storage profile supports encryption or not.
	// Example: false
	SupportsEncryption bool `json:"supportsEncryption,omitempty"`

	// A list of tags that represent the capabilities of this storage profile
	// Example: [ { \"key\" : \"tier\", \"value\": \"silver\" } ]
	Tags []*Tag `json:"tags"`

	// Date when the entity was last updated. The date is ISO 8601 and UTC.
	// Example: 2012-09-27
	UpdatedAt string `json:"updatedAt,omitempty"`

	// Indicates the type of volume associated with type of storage device.
	// Example: gp2 / io1 / sc1 / st1 / standard
	VolumeType string `json:"volumeType,omitempty"`
}


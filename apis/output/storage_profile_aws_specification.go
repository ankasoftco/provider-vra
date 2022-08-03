package models 
type StorageProfileAwsSpecification struct {

	// Indicates if a storage profile is default or not.
	// Example: true
	DefaultItem bool `json:"defaultItem,omitempty"`

	// A human-friendly description.
	Description string `json:"description,omitempty"`

	// Indicates the type of storage.
	// Example: ebs / instance-store
	// Required: true
	DeviceType *string `json:"deviceType"`

	// Indicates maximum I/O operations per second in range(1-20,000).
	// Example: 2000
	Iops string `json:"iops,omitempty"`

	// A human-friendly name used as an identifier in APIs that support this option.
	// Required: true
	Name *string `json:"name"`

	// A link to the region that is associated with the storage profile.
	// Example: 31186
	// Required: true
	RegionID *string `json:"regionId"`

	// Indicates whether this storage profile supports encryption or not.
	// Example: false
	SupportsEncryption bool `json:"supportsEncryption,omitempty"`

	// A list of tags that represent the capabilities of this storage profile
	// Example: [ { \"key\" : \"tier\", \"value\": \"silver\" } ]
	Tags []*Tag `json:"tags"`

	// Indicates the type of volume associated with type of storage.
	// Example: gp2 / io1 / sc1 / st1 / standard
	VolumeType string `json:"volumeType,omitempty"`
}


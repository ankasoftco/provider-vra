package models 
type StorageProfileGcpSpecification struct {

	// Indicates if a storage profile is default or not.
	// Example: true
	DefaultItem bool `json:"defaultItem,omitempty"`

	// A human-friendly description.
	Description string `json:"description,omitempty"`

	// A human-friendly name used as an identifier in APIs that support this option.
	// Required: true
	Name *string `json:"name"`

	// Indicates the type of disk.
	// Example: pd-standard / pd-ssd
	// Required: true
	PersistentDiskType *string `json:"persistentDiskType"`

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
}


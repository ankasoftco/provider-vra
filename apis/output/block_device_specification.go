package models 
type BlockDeviceSpecification struct {

	// Capacity of the block device in GB.
	// Example: 78
	// Required: true
	CapacityInGB *int32 `json:"capacityInGB"`

	// Constraints that are used to drive placement policies for the block device that is produced from this specification. Constraint expressions are matched against tags on existing placement targets.
	// Example: [ { \"mandatory\" : \"true\", \"expression\": \"environment\":\"prod\"}, {\"mandatory\" : \"false\", \"floor\": \"3rd\"} ]
	Constraints []*Constraint `json:"constraints"`

	// Additional custom properties that may be used to extend this resource.
	CustomProperties map[string]string `json:"customProperties,omitempty"`

	// The id of the deployment that is associated with this resource
	// Example: 123e4567-e89b-12d3-a456-426655440000
	DeploymentID string `json:"deploymentId,omitempty"`

	// A human-friendly description.
	Description string `json:"description,omitempty"`

	// Content of a disk, base64 encoded.
	// Example: dGVzdA
	DiskContentBase64 string `json:"diskContentBase64,omitempty"`

	// Indicates whether the block device should be encrypted or not.
	// Example: true
	Encrypted bool `json:"encrypted,omitempty"`

	// A human-friendly name used as an identifier in APIs that support this option.
	// Required: true
	Name *string `json:"name"`

	// Indicates whether the block device survives a delete action.
	// Example: true
	Persistent bool `json:"persistent,omitempty"`

	// The id of the project the current user belongs to.
	// Example: e058
	// Required: true
	ProjectID *string `json:"projectId"`

	// Reference to URI using which the block device has to be created.
	// Example: ami-0d4cfd66
	SourceReference string `json:"sourceReference,omitempty"`

	// A set of tag keys and optional values that should be set on any resource that is produced from this specification.
	// Example: [ { \"key\" : \"location\", \"value\": \"SOF\" } ]
	Tags []*Tag `json:"tags"`
}


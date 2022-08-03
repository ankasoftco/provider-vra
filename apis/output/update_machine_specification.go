package models 
type UpdateMachineSpecification struct {

	// Additional custom properties that may be used to extend the machine. Internal custom properties (for example, prefixed with: "__") are discarded.
	CustomProperties map[string]string `json:"customProperties,omitempty"`

	// Describes machine within the scope of your organization and is not propagated to the cloud
	Description string `json:"description,omitempty"`

	// A set of tag keys and optional values that should be set on any resource that is produced from this specification.
	// Example: [ { \"key\" : \"ownedBy\", \"value\": \"Rainpole\" } ]
	Tags []*Tag `json:"tags"`
}


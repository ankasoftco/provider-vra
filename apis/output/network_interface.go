package models 
type NetworkInterface struct {

	// HATEOAS of the entity
	// Required: true
	Links map[string]Href `json:"_links"`

	// A list of IP addresses allocated or in use by this network interface.
	// Example: [ \"10.1.2.190\" ]
	Addresses []string `json:"addresses"`

	// Set of ids of the cloud accounts this entity belongs to.
	// Example: [9e49]
	// Unique: true
	CloudAccountIds []string `json:"cloudAccountIds"`

	// Date when the entity was created. The date is in ISO 8601 and UTC.
	// Example: 2012-09-27
	CreatedAt string `json:"createdAt,omitempty"`

	// Additional properties that may be used to extend the base type.
	// Example: { \"awaitIp\" : \"true\" }
	CustomProperties map[string]string `json:"customProperties,omitempty"`

	// A human-friendly description.
	// Example: my-description
	Description string `json:"description,omitempty"`

	// The device index of this network interface.
	// Example: 1
	DeviceIndex int32 `json:"deviceIndex,omitempty"`

	// External entity Id on the provider side.
	// Example: i-cfe4-e241-e53b-756a9a2e25d2
	ExternalID string `json:"externalId,omitempty"`

	// The external regionId of the network interface.
	// Example: ap-northeast-2
	// Required: true
	ExternalRegionID *string `json:"externalRegionId"`

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

	// Email of the user that owns the entity.
	// Example: csp@vmware.com
	Owner string `json:"owner,omitempty"`

	// A list of security group ids this network interface is associated with
	SecurityGroupIds []string `json:"securityGroupIds"`

	// A set of tag keys and optional values that were set on this network interface.
	// Example: [ { \"key\" : \"vmware.enumeration.type\", \"value\": \"nec2_net_interface\" } ]
	Tags []*Tag `json:"tags"`

	// Date when the entity was last updated. The date is ISO 8601 and UTC.
	// Example: 2012-09-27
	UpdatedAt string `json:"updatedAt,omitempty"`
}


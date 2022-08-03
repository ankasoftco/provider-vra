package models 
type Machine struct {

	// HATEOAS of the entity
	// Required: true
	Links map[string]Href `json:"_links"`

	// Primary address allocated or in use by this machine. The actual type of the address depends on the adapter type. Typically it is either the public or the external IP address.
	// Example: 34.242.21.5
	Address string `json:"address,omitempty"`

	// The cloud config data in json-escaped yaml syntax
	BootConfig *MachineBootConfig `json:"bootConfig,omitempty"`

	// Set of ids of the cloud accounts this resource belongs to.
	// Example: [9e49]
	// Unique: true
	CloudAccountIds []string `json:"cloudAccountIds"`

	// Date when the entity was created. The date is in ISO 8601 and UTC.
	// Example: 2012-09-27
	CreatedAt string `json:"createdAt,omitempty"`

	// Additional properties that may be used to extend the base resource.
	// Example: { \"property\" : \"value\" }
	CustomProperties map[string]string `json:"customProperties,omitempty"`

	// Deployment id that is associated with this resource.
	// Example: 123e4567-e89b-12d3-a456-426655440000
	DeploymentID string `json:"deploymentId,omitempty"`

	// A human-friendly description.
	// Example: my-description
	Description string `json:"description,omitempty"`

	// External entity Id on the provider side.
	// Example: i-cfe4-e241-e53b-756a9a2e25d2
	ExternalID string `json:"externalId,omitempty"`

	// The external regionId of the resource.
	// Example: us-east-1
	// Required: true
	ExternalRegionID *string `json:"externalRegionId"`

	// The external zoneId of the resource.
	// Example: us-east-1a
	// Required: true
	ExternalZoneID *string `json:"externalZoneId"`

	// Hostname associated with this machine instance.
	Hostname string `json:"hostname,omitempty"`

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

	// Power state of machine.
	// Example: ON, OFF
	// Required: true
	// Enum: [ON OFF GUEST_OFF UNKNOWN SUSPEND]
	PowerState *string `json:"powerState"`

	// The id of the project this resource belongs to.
	// Example: 9e49
	ProjectID string `json:"projectId,omitempty"`

	// The provisioning status of the resource. One of three provisioning statuses.
	// `PROVISIONING`: The resource is being provisioned.
	// `READY`: The resource is already provisioned.
	// `SUSPEND`: The resource is being destroyed.
	//
	ProvisioningStatus string `json:"provisioningStatus,omitempty"`

	// Salt configuration on the machine
	SaltConfiguration *SaltConfiguration `json:"saltConfiguration,omitempty"`

	// A set of tag keys and optional values that were set on this resource.
	// Example: [ { \"key\" : \"ownedBy\", \"value\": \"Rainpole\" } ]
	Tags []*Tag `json:"tags"`

	// Date when the entity was last updated. The date is ISO 8601 and UTC.
	// Example: 2012-09-27
	UpdatedAt string `json:"updatedAt,omitempty"`
}


package models 
type ImageMappingDescription struct {

	// HATEOAS of the entity
	// Required: true
	Links map[string]Href `json:"_links"`

	// Set of ids of the cloud accounts this entity belongs to.
	// Example: [9e49]
	// Unique: true
	CloudAccountIds []string `json:"cloudAccountIds"`

	// Cloud config for this image. This cloud config will be merged during provisioning with other cloud configurations such as the bootConfig provided in MachineSpecification.
	// Example: runcmd:\n  - [\"mkdir\", \"/imageFolder\"]
	CloudConfig string `json:"cloudConfig,omitempty"`

	// Constraints that are used to drive placement policies for the image that is produced from this mapping.Constraint expressions are matched against tags on existing placement targets.
	// Example: [{\"mandatory\" : \"true\", \"expression\": \"environment\":\"prod\"}, {\"mandatory\" : \"false\", \"expression\": \"pci\"}]
	Constraints []*Constraint `json:"constraints"`

	// Date when the entity was created. The date is in ISO 8601 and UTC.
	// Example: 2012-09-27
	CreatedAt string `json:"createdAt,omitempty"`

	// Additional properties that may be used to extend the base type.
	// Example: { \"prop1\" : \"value1\" }
	CustomProperties map[string]string `json:"customProperties,omitempty"`

	// A human-friendly description.
	// Example: my-description
	Description string `json:"description,omitempty"`

	// External entity Id on the provider side.
	// Example: i-cfe4-e241-e53b-756a9a2e25d2
	ExternalID string `json:"externalId,omitempty"`

	// The regionId of the image
	// Example: us-east-1
	ExternalRegionID string `json:"externalRegionId,omitempty"`

	// The id of this resource instance
	// Example: 9e49
	// Required: true
	ID *string `json:"id"`

	// Indicates whether this fabric image is private. For vSphere, private images are considered to be templates and snapshots and public are Content Library Items
	// Example: true
	IsPrivate bool `json:"isPrivate,omitempty"`

	// A human-friendly name used as an identifier in APIs that support this option.
	// Example: my-name
	Name string `json:"name,omitempty"`

	// The id of the organization this entity belongs to.
	// Example: 42413b31-1716-477e-9a88-9dc1c3cb1cdf
	OrgID string `json:"orgId,omitempty"`

	// Operating System family of the image.
	// Example: linux, win
	OsFamily string `json:"osFamily,omitempty"`

	// Email of the user that owns the entity.
	// Example: csp@vmware.com
	Owner string `json:"owner,omitempty"`

	// Date when the entity was last updated. The date is ISO 8601 and UTC.
	// Example: 2012-09-27
	UpdatedAt string `json:"updatedAt,omitempty"`
}


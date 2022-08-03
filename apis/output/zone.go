package models 
type Zone struct {

	// HATEOAS of the entity
	// Required: true
	Links map[string]Href `json:"_links"`

	// Cloud account this zone belongs to.
	// Example: 9e49
	CloudAccountID string `json:"cloudAccountId,omitempty"`

	// Date when the entity was created. The date is in ISO 8601 and UTC.
	// Example: 2012-09-27
	CreatedAt string `json:"createdAt,omitempty"`

	// A list of key value pair of properties that will  be used
	// Example: {\"__isDefaultPlacementZone\": \"true\"}
	CustomProperties map[string]string `json:"customProperties,omitempty"`

	// A human-friendly description.
	// Example: my-description
	Description string `json:"description,omitempty"`

	// The id of the region for which this zone is defined
	// Example: us-east-1
	ExternalRegionID string `json:"externalRegionId,omitempty"`

	// The folder relative path to the datacenter where resources are deployed to. (only applicable for vSphere cloud zones)
	// Example: test-folder
	Folder string `json:"folder,omitempty"`

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

	// The placement policy for the zone.
	// Example: DEFAULT, SPREAD, BINPACK
	PlacementPolicy string `json:"placementPolicy,omitempty"`

	// A set of tag keys and optional values that were set on this placement.
	// Example: [ { \"key\" : \"dev\", \"value\": \" \" } ]
	Tags []*Tag `json:"tags"`

	// A set of tag keys and optional values for compute resource filtering.
	// Example: [ { \"key\" : \"compliance\", \"value\": \"pci\" } ]
	TagsToMatch []*Tag `json:"tagsToMatch"`

	// Date when the entity was last updated. The date is ISO 8601 and UTC.
	// Example: 2012-09-27
	UpdatedAt string `json:"updatedAt,omitempty"`
}


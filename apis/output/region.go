package models 
type Region struct {

	// HATEOAS of the entity
	// Required: true
	Links map[string]Href `json:"_links"`

	// The id of the cloud account this region belongs to.
	// Example: 9e49
	CloudAccountID string `json:"cloudAccountId,omitempty"`

	// Date when the entity was created. The date is in ISO 8601 and UTC.
	// Example: 2012-09-27
	CreatedAt string `json:"createdAt,omitempty"`

	// Unique identifier of region on the provider side.
	// Example: us-west
	// Required: true
	ExternalRegionID *string `json:"externalRegionId"`

	// The id of this resource instance
	// Example: 9e49
	// Required: true
	ID *string `json:"id"`

	// Name of region on the provider side. In vSphere, the name of the region is different from its id.
	// Example: us-west
	Name string `json:"name,omitempty"`

	// The id of the organization this entity belongs to.
	// Example: 42413b31-1716-477e-9a88-9dc1c3cb1cdf
	OrgID string `json:"orgId,omitempty"`

	// Email of the user that owns the entity.
	// Example: csp@vmware.com
	Owner string `json:"owner,omitempty"`

	// Date when the entity was last updated. The date is ISO 8601 and UTC.
	// Example: 2012-09-27
	UpdatedAt string `json:"updatedAt,omitempty"`
}


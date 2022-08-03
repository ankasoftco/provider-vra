package models 
type CloudAccountGcp struct {

	// HATEOAS of the entity
	// Required: true
	Links map[string]Href `json:"_links"`

	// GCP Client email
	// Example: 321743978432-compute@developer.gserviceaccount.com
	// Required: true
	ClientEmail *string `json:"clientEmail"`

	// Date when the entity was created. The date is in ISO 8601 and UTC.
	// Example: 2012-09-27
	CreatedAt string `json:"createdAt,omitempty"`

	// Additional properties that may be used to extend the base type.
	// Example: { \"isExternal\" : \"false\" }
	CustomProperties map[string]string `json:"customProperties,omitempty"`

	// A human-friendly description.
	// Example: my-description
	Description string `json:"description,omitempty"`

	// A list of regions that are enabled for provisioning on this cloud account
	EnabledRegions []*Region `json:"enabledRegions"`

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

	// GCP Private key ID
	// Example: 027f73d50a19452eedf5775a9b42c5083678abdf
	// Required: true
	PrivateKeyID *string `json:"privateKeyId"`

	// GCP Project ID
	// Example: example-gcp-project
	// Required: true
	ProjectID *string `json:"projectId"`

	// A set of tag keys and optional values that were set on on the Cloud Account
	// Example: [ { \"key\" : \"env\", \"value\": \"dev\" } ]
	Tags []*Tag `json:"tags"`

	// Date when the entity was last updated. The date is ISO 8601 and UTC.
	// Example: 2012-09-27
	UpdatedAt string `json:"updatedAt,omitempty"`
}


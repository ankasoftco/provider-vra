package models 
type StorageProfile struct {

	// HATEOAS of the entity
	// Required: true
	Links map[string]Href `json:"_links"`

	// Id of the cloud account this storage profile belongs to.
	// Example: [9e49]
	CloudAccountID string `json:"cloudAccountId,omitempty"`

	// Date when the entity was created. The date is in ISO 8601 and UTC.
	// Example: 2012-09-27
	CreatedAt string `json:"createdAt,omitempty"`

	// Indicates if a storage profile is default profile or not.
	// Required: true
	DefaultItem *bool `json:"defaultItem"`

	// A human-friendly description.
	// Example: my-description
	Description string `json:"description,omitempty"`

	// Map of storage properties that are to be applied on disk while provisioning.
	// Example: { \"diskProperties\": {\n                    \"provisioningType\": \"thin\",\n                    \"sharesLevel\": \"low\",\n                    \"shares\": \"500\",\n                    \"limitIops\": \"500\",\n                    \"diskType\": \"firstClass\"\n                } }
	DiskProperties map[string]string `json:"diskProperties,omitempty"`

	// The id of the region for which this profile is defined
	// Example: us-east-1
	ExternalRegionID string `json:"externalRegionId,omitempty"`

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

	// Indicates whether this storage profile supports encryption or not.
	SupportsEncryption bool `json:"supportsEncryption,omitempty"`

	// A list of tags that represent the capabilities of this storage profile
	// Example: [ { \"key\" : \"tier\", \"value\": \"silver\" } ]
	Tags []*Tag `json:"tags"`

	// Date when the entity was last updated. The date is ISO 8601 and UTC.
	// Example: 2012-09-27
	UpdatedAt string `json:"updatedAt,omitempty"`
}


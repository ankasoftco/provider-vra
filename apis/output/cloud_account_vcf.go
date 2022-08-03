package models 
type CloudAccountVcf struct {

	// HATEOAS of the entity
	// Required: true
	Links map[string]Href `json:"_links"`

	// Create default cloud zones for the enabled regions.
	// Example: true
	CreateDefaultZones bool `json:"createDefaultZones,omitempty"`

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

	// NSX Hostname in a workload domain
	// Example: nsx.mycompany.com
	// Required: true
	NsxHostName *string `json:"nsxHostName"`

	// Username to authenticate to NSX manager in workload domain
	// Example: administrator@mycompany.com
	// Required: true
	NsxUsername *string `json:"nsxUsername"`

	// The id of the organization this entity belongs to.
	// Example: 42413b31-1716-477e-9a88-9dc1c3cb1cdf
	OrgID string `json:"orgId,omitempty"`

	// Email of the user that owns the entity.
	// Example: csp@vmware.com
	Owner string `json:"owner,omitempty"`

	// SDDC Manager ID
	// Example: 2e5bb71d-0c14-4066-a999-2cb6c693654a
	// Required: true
	SddcManagerID *string `json:"sddcManagerId"`

	// A set of tag keys and optional values to set on the Cloud Account.Cloud account capability tags may enable different features.
	// Example: [ { \"key\" : \"env\", \"value\": \"dev\" } ]
	Tags []*Tag `json:"tags"`

	// Date when the entity was last updated. The date is ISO 8601 and UTC.
	// Example: 2012-09-27
	UpdatedAt string `json:"updatedAt,omitempty"`

	// vCenter Hostname in a workload domain
	// Example: vcenter.mycompany.com
	// Required: true
	VcenterHostName *string `json:"vcenterHostName"`

	// Username to authenticate to vCenter Server in workload domain
	// Example: administrator@mycompany.com
	// Required: true
	VcenterUsername *string `json:"vcenterUsername"`

	// Id of the VCF worload domain.
	// Example: 587db412-6037-43e4-8e1e-49ebbaf6cd35
	// Required: true
	VcfDomainID *string `json:"vcfDomainId"`

	// Name of the VCF worload domain.
	// Example: Workload Domain - 1
	// Required: true
	VcfDomainName *string `json:"vcfDomainName"`
}


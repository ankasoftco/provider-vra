package models 
type FabricNetworkVsphere struct {

	// HATEOAS of the entity
	// Required: true
	Links map[string]Href `json:"_links"`

	// Network CIDR to be used.
	// Example: 10.1.2.0/24
	Cidr string `json:"cidr,omitempty"`

	// Set of ids of the cloud accounts this entity belongs to.
	// Example: [9e49]
	// Unique: true
	CloudAccountIds []string `json:"cloudAccountIds"`

	// Date when the entity was created. The date is in ISO 8601 and UTC.
	// Example: 2012-09-27
	CreatedAt string `json:"createdAt,omitempty"`

	// Custom properties of the fabric network instance
	CustomProperties map[string]string `json:"customProperties,omitempty"`

	// IPv4 default gateway to be used.
	// Example: 10.1.2.1
	DefaultGateway string `json:"defaultGateway,omitempty"`

	// IPv6 default gateway to be used.
	// Example: 2001:eeee:6bd:2a::1
	DefaultIPV6Gateway string `json:"defaultIpv6Gateway,omitempty"`

	// A human-friendly description.
	// Example: my-description
	Description string `json:"description,omitempty"`

	// A list of DNS search domains that were set on this resource instance.
	// Example: [vmware.com]
	DNSSearchDomains []string `json:"dnsSearchDomains"`

	// A list of DNS server addresses that were set on this resource instance.
	// Example: [1.1.1.1]
	DNSServerAddresses []string `json:"dnsServerAddresses"`

	// Domain value.
	// Example: sqa.local
	Domain string `json:"domain,omitempty"`

	// External entity Id on the provider side.
	// Example: i-cfe4-e241-e53b-756a9a2e25d2
	ExternalID string `json:"externalId,omitempty"`

	// The id of the region for which this network is defined
	// Example: us-east-1
	ExternalRegionID string `json:"externalRegionId,omitempty"`

	// The id of this resource instance
	// Example: 9e49
	// Required: true
	ID *string `json:"id"`

	// Network IPv6 CIDR to be used.
	// Example: 2001:eeee:6bd:2a::1/64
	IPV6Cidr string `json:"ipv6Cidr,omitempty"`

	// Indicates whether this is the default subnet for the zone.
	IsDefault bool `json:"isDefault,omitempty"`

	// Indicates whether the sub-network supports public IP assignment.
	IsPublic bool `json:"isPublic,omitempty"`

	// A human-friendly name used as an identifier in APIs that support this option.
	// Example: my-name
	Name string `json:"name,omitempty"`

	// The id of the organization this entity belongs to.
	// Example: 42413b31-1716-477e-9a88-9dc1c3cb1cdf
	OrgID string `json:"orgId,omitempty"`

	// Email of the user that owns the entity.
	// Example: csp@vmware.com
	Owner string `json:"owner,omitempty"`

	// A set of tag keys and optional values that were set on this resource instance.
	// Example: [ { \"key\" : \"fast-network\", \"value\": \"true\" } ]
	Tags []*Tag `json:"tags"`

	// Date when the entity was last updated. The date is ISO 8601 and UTC.
	// Example: 2012-09-27
	UpdatedAt string `json:"updatedAt,omitempty"`
}


package models 
type ExternalNetworkIPRange struct {

	// HATEOAS of the entity
	// Required: true
	Links map[string]Href `json:"_links"`

	// Address space that the range belongs to
	AddressSpaceID string `json:"addressSpaceId,omitempty"`

	// Date when the entity was created. The date is in ISO 8601 and UTC.
	// Example: 2012-09-27
	CreatedAt string `json:"createdAt,omitempty"`

	// A human-friendly description.
	// Example: my-description
	Description string `json:"description,omitempty"`

	// DNS domain search (in order)
	DNSSearchDomains []string `json:"dnsSearchDomains"`

	// DNS IP addresses of the range.
	DNSServerAddresses []string `json:"dnsServerAddresses"`

	// DNS domain of the range.
	Domain string `json:"domain,omitempty"`

	// End IP address of the range.
	// Required: true
	EndIPAddress *string `json:"endIPAddress"`

	// External entity Id on the provider side.
	// Example: i-cfe4-e241-e53b-756a9a2e25d2
	ExternalID string `json:"externalId,omitempty"`

	// The gateway address of the range
	GatewayAddress string `json:"gatewayAddress,omitempty"`

	// The id of this resource instance
	// Example: 9e49
	// Required: true
	ID *string `json:"id"`

	// IP address version: IPv4 or IPv6. Default: IPv4.
	// Enum: [IPv4 IPv6]
	IPVersion string `json:"ipVersion,omitempty"`

	// A human-friendly name used as an identifier in APIs that support this option.
	// Example: my-name
	Name string `json:"name,omitempty"`

	// The id of the organization this entity belongs to.
	// Example: 42413b31-1716-477e-9a88-9dc1c3cb1cdf
	OrgID string `json:"orgId,omitempty"`

	// Email of the user that owns the entity.
	// Example: csp@vmware.com
	Owner string `json:"owner,omitempty"`

	// Start IP address of the range.
	// Required: true
	StartIPAddress *string `json:"startIPAddress"`

	// Subnet prefix length (synonymous with "netmask")
	// Required: true
	SubnetPrefixLength *int32 `json:"subnetPrefixLength"`

	// A set of tag keys and optional values that were set on this resource instance.
	// Example: [ { \"key\" : \"ipv6-range\", \"value\": \"true\" } ]
	Tags []*Tag `json:"tags"`

	// Date when the entity was last updated. The date is ISO 8601 and UTC.
	// Example: 2012-09-27
	UpdatedAt string `json:"updatedAt,omitempty"`
}


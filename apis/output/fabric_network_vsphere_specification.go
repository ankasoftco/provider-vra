package models 
type FabricNetworkVsphereSpecification struct {

	// Network CIDR to be used.
	// Example: 10.1.2.0/24
	Cidr string `json:"cidr,omitempty"`

	// IPv4 default gateway to be used.
	// Example: 10.1.2.1
	DefaultGateway string `json:"defaultGateway,omitempty"`

	// IPv6 default gateway to be used.
	// Example: 2001:eeee:6bd:2a::1
	DefaultIPV6Gateway string `json:"defaultIpv6Gateway,omitempty"`

	// A list of DNS search domains that were set on this resource instance.
	// Example: [vmware.com]
	DNSSearchDomains []string `json:"dnsSearchDomains"`

	// A list of DNS server addresses that were set on this resource instance.
	// Example: [1.1.1.1]
	DNSServerAddresses []string `json:"dnsServerAddresses"`

	// Domain value.
	// Example: sqa.local
	Domain string `json:"domain,omitempty"`

	// Network IPv6 CIDR to be used.
	// Example: 2001:eeee:6bd:2a::1/64
	IPV6Cidr string `json:"ipv6Cidr,omitempty"`

	// Indicates whether this is the default subnet for the zone.
	IsDefault *bool `json:"isDefault,omitempty"`

	// Indicates whether the sub-network supports public IP assignment.
	IsPublic *bool `json:"isPublic,omitempty"`

	// A set of tag keys and optional values that were set on this resource instance.
	// Example: [ { \"key\" : \"fast-network\", \"value\": \"true\" } ]
	Tags []*Tag `json:"tags"`
}


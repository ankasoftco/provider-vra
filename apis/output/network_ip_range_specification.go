package models 
type NetworkIPRangeSpecification struct {

	// A human-friendly description.
	Description string `json:"description,omitempty"`

	// End IP address of the range.
	// Required: true
	EndIPAddress *string `json:"endIPAddress"`

	// The Ids of the fabric networks.
	// Unique: true
	FabricNetworkIds []string `json:"fabricNetworkIds"`

	// IP address version: IPv4 or IPv6. Default: IPv4.
	// Enum: [IPv4 IPv6]
	IPVersion string `json:"ipVersion,omitempty"`

	// A human-friendly name used as an identifier in APIs that support this option.
	// Required: true
	Name *string `json:"name"`

	// Start IP address of the range.
	// Required: true
	StartIPAddress *string `json:"startIPAddress"`

	// A set of tag keys and optional values that were set on this resource instance.
	// Example: [ { \"key\" : \"fast-network\", \"value\": \"true\" } ]
	Tags []*Tag `json:"tags"`
}


package models 
type Rule struct {

	// Type of access (Allow, Deny or Drop) for the security rule. Allow is default. Traffic that does not match any rules will be denied.
	// Example: Allow
	// Required: true
	// Enum: [Allow Deny Drop]
	Access *string `json:"access"`

	// Direction of the security rule (inbound or outboud).
	// Example: Outbound
	// Required: true
	// Enum: [Inbound Outbound]
	Direction *string `json:"direction"`

	// IP address(es) in CIDR format which the security rule applies to.
	// Example: 66.170.99.2/32
	// Required: true
	IPRangeCidr *string `json:"ipRangeCidr"`

	// Name of security rule.
	// Example: 5756f7e2
	Name string `json:"name,omitempty"`

	// Ports the security rule applies to.
	// Example: 443, 1-655535
	// Required: true
	Ports *string `json:"ports"`

	// Protocol the security rule applies to.
	// Example: ANY, TCP, UDP
	Protocol string `json:"protocol,omitempty"`

	// Service defined by the provider (such as: SSH, HTTPS). Either service or protocol have to be specified.
	// Example: HTTPS, SSH
	Service string `json:"service,omitempty"`
}


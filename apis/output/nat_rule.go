package models 
type NatRule struct {

	// Description of the NAT rule.
	Description string `json:"description,omitempty"`

	// The external IP address of the outbound or routed network
	// Example: any, 10.20.156.101
	// Read Only: true
	DestinationAddress string `json:"destinationAddress,omitempty"`

	// The edge gateway port. Default is `any`
	// Example: any, 80, 5000-5100
	DestinationPorts string `json:"destinationPorts,omitempty"`

	// Index in which the rule must be applied
	// Required: true
	Index *int32 `json:"index"`

	// Kind of NAT: NAT44/NAT64/NAT66. Only NAT44 is supported currently and it is the default value.
	// Example: NAT44
	Kind string `json:"kind,omitempty"`

	// The protocol of the incoming requests. Default is TCP.
	// Example: TCP, UDP
	Protocol string `json:"protocol,omitempty"`

	// Unique ID of the NAT Rule
	// Read Only: true
	RuleID string `json:"ruleId,omitempty"`

	// The IP of the external source. Default is `any`
	// Example: any, 10.20.156.101
	SourceIPs string `json:"sourceIPs,omitempty"`

	// Ports from where the request is originating. Default is `any`
	// Example: any, 80, 5000-5100
	SourcePorts string `json:"sourcePorts,omitempty"`

	// A links to target load balancer or a machine's network interface where the request will be forwarded.
	// Example: /iaas/api/load-balancers/try6-45ef, /iaas/api/machines/ht54-a472/network-interfaces/dyd6-d67e
	// Required: true
	TargetLink *string `json:"targetLink"`

	// The machine port where the request will be forwarded. Default is `any`
	// Example: any, 80, 5000-5100
	TranslatedPorts string `json:"translatedPorts,omitempty"`

	// Type of the NAT rule. Only DNAT is supported currently.
	// Example: DNAT
	Type string `json:"type,omitempty"`
}


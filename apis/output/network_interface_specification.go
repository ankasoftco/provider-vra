package models 
type NetworkInterfaceSpecification struct {

	// A list of IP addresses allocated or in use by this network interface.
	// Example: [ \"10.1.2.190\" ]
	Addresses []string `json:"addresses"`

	// Additional properties that may be used to extend the base type.
	// Example: { \"awaitIp\" : \"true\" }
	CustomProperties map[string]string `json:"customProperties,omitempty"`

	// A human-friendly description.
	Description string `json:"description,omitempty"`

	// The device index of this network interface.
	// Example: 1
	DeviceIndex int32 `json:"deviceIndex,omitempty"`

	// Id of the fabric network for the network interface. Either networkId or fabricNetworkId can be specified, not both.
	// Example: 54097407-4532-460c-94a8-8f9e18f4c925
	FabricNetworkID string `json:"fabricNetworkId,omitempty"`

	// MAC address of the network interface.
	// Example: [ \"00:50:56:99:d8:34\" ]
	MacAddress string `json:"macAddress,omitempty"`

	// A human-friendly name used as an identifier in APIs that support this option.
	Name string `json:"name,omitempty"`

	// Id of the network for the network interface. Either networkId or fabricNetworkId can be specified, not both.
	// Example: 54097407-4532-460c-94a8-8f9e18f4c925
	NetworkID string `json:"networkId,omitempty"`

	// A list of security group ids which this network interface will be assigned to.
	SecurityGroupIds []string `json:"securityGroupIds"`
}


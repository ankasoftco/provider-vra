package models 
type UpdateNetworkInterfaceSpecification struct {

	// Set IPv4 address for the machine network interface. The change will not propagate to cloud endpoint for provisioned machines.
	Address string `json:"address,omitempty"`

	// Additional custom properties that may be used to extend the machine. Internal custom properties (for example, prefixed with: "__") can not be updated.
	CustomProperties map[string]string `json:"customProperties,omitempty"`

	// Describes the network interface of the machine within the scope of your organization and is not propagated to the cloud
	Description string `json:"description,omitempty"`

	// Network interface name used during machine network interface provisioning. This property only takes effect if it is set before machine provisioning starts. The change will not propagate to cloud endpoint for provisioned machines.
	Name string `json:"name,omitempty"`
}


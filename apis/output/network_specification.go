package models 
type NetworkSpecification struct {

	// Constraints that are used to drive placement policies for the network that is produced from this specification, related with the network profile. Constraint expressions are matched against tags on existing placement targets.
	// Example: [{\"mandatory\" : \"true\", \"expression\": \"environment\":\"prod\"}, {\"mandatory\" : \"false\", \"expression\": \"pci\"}]
	Constraints []*Constraint `json:"constraints"`

	// Flag to indicate if the network creation should create a gateway. Default is true.
	// Example: true
	CreateGateway bool `json:"createGateway,omitempty"`

	// Additional custom properties that may be used to extend this resource.
	CustomProperties map[string]string `json:"customProperties,omitempty"`

	// The id of the deployment that is associated with this resource
	// Example: 123e4567-e89b-12d3-a456-426655440000
	DeploymentID string `json:"deploymentId,omitempty"`

	// A human-friendly description.
	Description string `json:"description,omitempty"`

	// A human-friendly name used as an identifier in APIs that support this option.
	// Required: true
	Name *string `json:"name"`

	// Flag to indicate if the network needs to have outbound access or not. Default is true. This field will be ignored if there is proper input for networkType customProperty
	// Example: true
	OutboundAccess bool `json:"outboundAccess,omitempty"`

	// The id of the project the current user belongs to.
	// Example: e058
	// Required: true
	ProjectID *string `json:"projectId"`

	// A set of tag keys and optional values that should be set on any resource that is produced from this specification.
	// Example: [ { \"key\" : \"vmware.enumeration.type\", \"value\": \"nec2_vpc\" } ]
	Tags []*Tag `json:"tags"`
}


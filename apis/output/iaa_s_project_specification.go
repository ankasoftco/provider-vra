package models 
type IaaSProjectSpecification struct {

	// List of administrator users associated with the project. Only administrators can manage project's configuration.
	// Example: [{ \"email\":\"administrator@vmware.com\" }]
	Administrators []*User `json:"administrators"`

	// List of storage, network and extensibility constraints to be applied when provisioning through this project.
	// Example: {\"network\":[{\"mandatory\": \"true\", \"expression\": \"env:dev\"}],\"storage\":[{\"mandatory\": \"false\", \"expression\": \"gold\"}],\"extensibility\":[{\"mandatory\": \"false\", \"expression\": \"key:value\"}]}
	Constraints map[string][]Constraint `json:"constraints,omitempty"`

	// The project custom properties which are added to all requests in this project
	// Example: { \"property\" : \"value\" }
	CustomProperties map[string]string `json:"customProperties,omitempty"`

	// A human-friendly description.
	Description string `json:"description,omitempty"`

	// The naming template to be used for machines provisioned in this project
	// Example: ${project.name}-test-${####}
	MachineNamingTemplate string `json:"machineNamingTemplate,omitempty"`

	// List of member users associated with the project.
	// Example: [{ \"email\":\"member@vmware.com\" }]
	Members []*User `json:"members"`

	// A human-friendly name used as an identifier in APIs that support this option.
	// Required: true
	Name *string `json:"name"`

	// The timeout that should be used for Blueprint operations and Provisioning tasks. The timeout is in seconds
	// Example: 30
	OperationTimeout *int64 `json:"operationTimeout,omitempty"`

	// Placement policy for the project. Determines how a zone will be selected for provisioning. DEFAULT or SPREAD.
	// Example: DEFAULT
	PlacementPolicy string `json:"placementPolicy,omitempty"`

	// Specifies whether the resources in this projects are shared or not. If not set default will be used.
	// Example: true
	SharedResources bool `json:"sharedResources,omitempty"`

	// List of viewer users associated with the project.
	// Example: [{ \"email\":\"viewer@vmware.com\" }]
	Viewers []*User `json:"viewers"`

	// List of configurations for zone assignment to a project.
	ZoneAssignmentConfigurations []*ZoneAssignmentSpecification `json:"zoneAssignmentConfigurations"`
}


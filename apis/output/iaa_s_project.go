package models 
type IaaSProject struct {

	// HATEOAS of the entity
	// Required: true
	Links map[string]Href `json:"_links"`

	// List of administrator users associated with the project. Only administrators can manage project's configuration.
	// Example: [ { \"email\":\"administrator@vmware.com\" } ]
	Administrators []*User `json:"administrators"`

	// List of storage, network and extensibility constraints to be applied when provisioning through this project.
	// Example: {\"network\":[{\"mandatory\": \"true\", \"expression\": \"env:dev\"}],\"storage\":[{\"mandatory\": \"false\", \"expression\": \"gold\"}],\"extensibility\":[{\"mandatory\": \"false\", \"expression\": \"key:value\"}]}
	Constraints map[string][]Constraint `json:"constraints,omitempty"`

	// Date when the entity was created. The date is in ISO 8601 and UTC.
	// Example: 2012-09-27
	CreatedAt string `json:"createdAt,omitempty"`

	// The project custom properties which are added to all requests in this project
	// Example: { \"property\" : \"value\" }
	CustomProperties map[string]string `json:"customProperties,omitempty"`

	// A human-friendly description.
	// Example: my-description
	Description string `json:"description,omitempty"`

	// The id of this resource instance
	// Example: 9e49
	// Required: true
	ID *string `json:"id"`

	// The naming template to be used for machines provisioned in this project
	// Example: ${project.name}-test-${####}
	MachineNamingTemplate string `json:"machineNamingTemplate,omitempty"`

	// List of member users associated with the project.
	// Example: [ { \"email\":\"member@vmware.com\" } ]
	Members []*User `json:"members"`

	// A human-friendly name used as an identifier in APIs that support this option.
	// Example: my-name
	Name string `json:"name,omitempty"`

	// The timeout that should be used for Blueprint operations and Provisioning tasks. The timeout is in seconds
	OperationTimeout int64 `json:"operationTimeout,omitempty"`

	// The id of the organization this entity belongs to.
	// Example: 42413b31-1716-477e-9a88-9dc1c3cb1cdf
	OrgID string `json:"orgId,omitempty"`

	// Email of the user that owns the entity.
	// Example: csp@vmware.com
	Owner string `json:"owner,omitempty"`

	// Placement policy for the project. Determines how a zone will be selected for provisioning. DEFAULT or SPREAD.
	// Example: DEFAULT
	PlacementPolicy string `json:"placementPolicy,omitempty"`

	// Specifies whether the resources in this projects are shared or not.
	SharedResources bool `json:"sharedResources,omitempty"`

	// Date when the entity was last updated. The date is ISO 8601 and UTC.
	// Example: 2012-09-27
	UpdatedAt string `json:"updatedAt,omitempty"`

	// List of viewer users associated with the project.
	// Example: [ { \"email\":\"viewer@vmware.com\" } ]
	Viewers []*User `json:"viewers"`

	// List of Cloud Zones assigned to this project. You can limit deployment to a single region or allow multi-region placement by adding more than one cloud zone to a project. A cloud zone lists available resources. Use tags on resources to control workload placement.
	Zones []*ZoneAssignment `json:"zones"`
}


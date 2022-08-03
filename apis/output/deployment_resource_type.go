package models 
type DeploymentResourceType struct {

	// Optional. Account type to which the resource type belongs to. Example: AWS, Azure etc
	// Enum: [AWS Azure GCP vSphere vSphere-cloud Azure-EA NSX-V NSX-T NSX-P NSX-P-cloud vCloud Director VMC Puppet Ansible]
	AccountType string `json:"accountType,omitempty"`

	// Composable with other types or not
	Composable bool `json:"composable,omitempty"`

	// Time at which the resource type was created.
	// Format: date-time
	CreatedAt strfmt.DateTime `json:"createdAt,omitempty"`

	// Name of the user who created the resource type.
	CreatedBy string `json:"createdBy,omitempty"`

	// Resource type description
	Description string `json:"description,omitempty"`

	// Resource display name
	DisplayName string `json:"displayName,omitempty"`

	// Resource type id
	ID string `json:"id,omitempty"`

	// Resource type name
	Name string `json:"name,omitempty"`

	// Org ID where resource type belongs
	OrgID string `json:"orgId,omitempty"`

	// Project ID where resource type belongs
	ProjectID string `json:"projectId,omitempty"`

	// Provider Id
	ProviderID string `json:"providerId,omitempty"`

	// Json schema that represents resource type, a simplified version of http://json-schema.org/latest/json-schema-validation.html#rfc.section.5
	Schema interface{} `json:"schema,omitempty"`

	// Time at which the resource type was updated.
	// Format: date-time
	UpdatedAt strfmt.DateTime `json:"updatedAt,omitempty"`

	// Name of the user who updated the resource type.
	UpdatedBy string `json:"updatedBy,omitempty"`
}


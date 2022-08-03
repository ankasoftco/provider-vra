package models 
type CustomNaming struct {

	// HATEOAS of the entity
	// Required: true
	Links map[string]Href `json:"_links"`

	// Date when the entity was created. The date is in ISO 8601 and UTC.
	// Example: 2012-09-27
	CreatedAt string `json:"createdAt,omitempty"`

	// A human-friendly description.
	// Example: my-description
	Description string `json:"description,omitempty"`

	// The id of this resource instance
	// Example: 9e49
	// Required: true
	ID *string `json:"id"`

	// A human-friendly name used as an identifier in APIs that support this option.
	// Example: my-name
	Name string `json:"name,omitempty"`

	// The id of the organization this entity belongs to.
	// Example: 42413b31-1716-477e-9a88-9dc1c3cb1cdf
	OrgID string `json:"orgId,omitempty"`

	// Email of the user that owns the entity.
	// Example: csp@vmware.com
	Owner string `json:"owner,omitempty"`

	// Set of projects associated with custom name
	// Example: [{\"defaultOrg\":true,\"active\":true,\"id\":\"3fa85f64-5717-4562-b3fc-2c963f66afa6\",\"projectName\":\"string\",\"projectId\":\"string\",\"orgId\":\"string\"}]
	// Unique: true
	Projects []*CustomNamingProject `json:"projects"`

	// Set of templates associated with custom name
	// Example: [{\"uniqueName\":true,\"staticPattern\":\"string\",\"counters\":[{\"currentCounter\":0,\"cnResourceType\":\"COMPUTE\",\"active\":true,\"id\":\"3fa85f64-5717-4562-b3fc-2c963f66afa6\",\"projectId\":\"string\"}],\"incrementStep\":0,\"pattern\":\"string\",\"startCounter\":0,\"id\":\"3fa85f64-5717-4562-b3fc-2c963f66afa6\",\"resourceTypeName\":\"string\",\"resourceType\":\"COMPUTE\",\"resourceDefault\":true}]
	// Unique: true
	Templates []*CustomNamingTemplate `json:"templates"`

	// Date when the entity was last updated. The date is ISO 8601 and UTC.
	// Example: 2012-09-27
	UpdatedAt string `json:"updatedAt,omitempty"`
}


package models 
type K8SLimitRange struct {

	// active
	Active bool `json:"active,omitempty"`

	// cluster Id
	// Format: uuid
	ClusterID strfmt.UUID `json:"clusterId,omitempty"`

	// created millis
	CreatedMillis int64 `json:"createdMillis,omitempty"`

	// definition
	Definition string `json:"definition,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// id
	// Format: uuid
	ID strfmt.UUID `json:"id,omitempty"`

	// managed
	Managed bool `json:"managed,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// namespace Id
	// Format: uuid
	NamespaceID strfmt.UUID `json:"namespaceId,omitempty"`

	// namespace name
	NamespaceName string `json:"namespaceName,omitempty"`

	// org Id
	OrgID string `json:"orgId,omitempty"`

	// resource
	Resource string `json:"resource,omitempty"`

	// type
	Type string `json:"type,omitempty"`

	// updated millis
	UpdatedMillis int64 `json:"updatedMillis,omitempty"`
}


package models 
type K8SNamespace struct {

	// auth credentials link
	AuthCredentialsLink string `json:"authCredentialsLink,omitempty"`

	// cluster Id
	// Format: uuid
	ClusterID strfmt.UUID `json:"clusterId,omitempty"`

	// content
	Content string `json:"content,omitempty"`

	// created millis
	CreatedMillis int64 `json:"createdMillis,omitempty"`

	// custom properties
	CustomProperties interface{} `json:"customProperties,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// id
	// Format: uuid
	ID strfmt.UUID `json:"id,omitempty"`

	// installer Id
	// Format: uuid
	InstallerID strfmt.UUID `json:"installerId,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// org Id
	OrgID string `json:"orgId,omitempty"`

	// owner
	Owner string `json:"owner,omitempty"`

	// project Id
	ProjectID string `json:"projectId,omitempty"`

	// registered
	Registered bool `json:"registered,omitempty"`

	// shared
	Shared bool `json:"shared,omitempty"`

	// status
	// Enum: [ALLOCATED READY FAILED UNREACHABLE TERMINATING REMOVED]
	Status string `json:"status,omitempty"`

	// updated millis
	UpdatedMillis int64 `json:"updatedMillis,omitempty"`

	// zone project assignment Id
	// Format: uuid
	ZoneProjectAssignmentID strfmt.UUID `json:"zoneProjectAssignmentId,omitempty"`
}


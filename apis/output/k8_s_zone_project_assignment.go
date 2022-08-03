package models 
type K8SZoneProjectAssignment struct {

	// created millis
	CreatedMillis int64 `json:"createdMillis,omitempty"`

	// id
	// Format: uuid
	ID strfmt.UUID `json:"id,omitempty"`

	// max number namespaces
	MaxNumberNamespaces int32 `json:"maxNumberNamespaces,omitempty"`

	// max number supervisor namespaces
	MaxNumberSupervisorNamespaces int32 `json:"maxNumberSupervisorNamespaces,omitempty"`

	// org Id
	OrgID string `json:"orgId,omitempty"`

	// priority
	Priority int32 `json:"priority,omitempty"`

	// project Id
	ProjectID string `json:"projectId,omitempty"`

	// updated millis
	UpdatedMillis int64 `json:"updatedMillis,omitempty"`
}


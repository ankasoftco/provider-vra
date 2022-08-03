package models 
type K8SZoneResourceAssignment struct {

	// created millis
	CreatedMillis int64 `json:"createdMillis,omitempty"`

	// enabled
	Enabled bool `json:"enabled,omitempty"`

	// id
	// Format: uuid
	ID strfmt.UUID `json:"id,omitempty"`

	// org Id
	OrgID string `json:"orgId,omitempty"`

	// priority
	Priority int32 `json:"priority,omitempty"`

	// resource Id
	ResourceID string `json:"resourceId,omitempty"`

	// resource type
	// Enum: [PKS_PLAN]
	ResourceType string `json:"resourceType,omitempty"`

	// tag links
	TagLinks []string `json:"tagLinks"`

	// tags
	Tags []*TagState `json:"tags"`

	// updated millis
	UpdatedMillis int64 `json:"updatedMillis,omitempty"`
}


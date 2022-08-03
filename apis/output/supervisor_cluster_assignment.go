package models 
type SupervisorClusterAssignment struct {

	// cluster self link Id
	ClusterSelfLinkID string `json:"clusterSelfLinkId,omitempty"`

	// created millis
	CreatedMillis int64 `json:"createdMillis,omitempty"`

	// id
	// Format: uuid
	ID strfmt.UUID `json:"id,omitempty"`

	// org Id
	OrgID string `json:"orgId,omitempty"`

	// priority
	Priority int32 `json:"priority,omitempty"`

	// tag links
	TagLinks []string `json:"tagLinks"`

	// tags
	Tags []*TagState `json:"tags"`

	// updated millis
	UpdatedMillis int64 `json:"updatedMillis,omitempty"`
}


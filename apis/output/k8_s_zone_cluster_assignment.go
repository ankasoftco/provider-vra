package models 
type K8SZoneClusterAssignment struct {

	// cluster Id
	// Format: uuid
	ClusterID strfmt.UUID `json:"clusterId,omitempty"`

	// created millis
	CreatedMillis int64 `json:"createdMillis,omitempty"`

	// id
	// Format: uuid
	ID strfmt.UUID `json:"id,omitempty"`

	// org Id
	OrgID string `json:"orgId,omitempty"`

	// pks cluster Uuid
	PksClusterUUID string `json:"pksClusterUuid,omitempty"`

	// prefer master Ip
	PreferMasterIP bool `json:"preferMasterIp,omitempty"`

	// priority
	Priority int32 `json:"priority,omitempty"`

	// tag links
	TagLinks []string `json:"tagLinks"`

	// tags
	Tags []*TagState `json:"tags"`

	// updated millis
	UpdatedMillis int64 `json:"updatedMillis,omitempty"`
}


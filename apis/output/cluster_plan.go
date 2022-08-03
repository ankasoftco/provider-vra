package models 
type ClusterPlan struct {

	// cloud account self link Id
	CloudAccountSelfLinkID string `json:"cloudAccountSelfLinkId,omitempty"`

	// created millis
	CreatedMillis int64 `json:"createdMillis,omitempty"`

	// definition
	Definition interface{} `json:"definition,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// id
	// Format: uuid
	ID strfmt.UUID `json:"id,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// org Id
	OrgID string `json:"orgId,omitempty"`

	// type
	// Enum: [TANZU_CLUSTER_PLAN]
	Type string `json:"type,omitempty"`

	// updated millis
	UpdatedMillis int64 `json:"updatedMillis,omitempty"`
}


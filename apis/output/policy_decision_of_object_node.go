package models 
type PolicyDecisionOfObjectNode struct {

	// description
	Description string `json:"description,omitempty"`

	// dry run Id
	// Format: uuid
	DryRunID strfmt.UUID `json:"dryRunId,omitempty"`

	// dry run sub task Id
	// Format: uuid
	DryRunSubTaskID strfmt.UUID `json:"dryRunSubTaskId,omitempty"`

	// effective policy definition
	EffectivePolicyDefinition interface{} `json:"effectivePolicyDefinition,omitempty"`

	// id
	// Format: uuid
	ID strfmt.UUID `json:"id,omitempty"`

	// org Id
	OrgID string `json:"orgId,omitempty"`

	// policies
	Policies []*PolicyDecisionPolicy `json:"policies"`

	// project Id
	ProjectID string `json:"projectId,omitempty"`

	// target Id
	TargetID string `json:"targetId,omitempty"`

	// target name
	TargetName string `json:"targetName,omitempty"`

	// timestamp
	// Format: date-time
	Timestamp strfmt.DateTime `json:"timestamp,omitempty"`

	// type Id
	TypeID string `json:"typeId,omitempty"`
}


package models 
type Policy struct {

	// Policy creation timestamp.
	// Format: date-time
	CreatedAt strfmt.DateTime `json:"createdAt,omitempty"`

	// Policy author.
	CreatedBy string `json:"createdBy,omitempty"`

	// Policy-type-specific target object filter criteria applied during enforcement.
	Criteria *Criteria `json:"criteria,omitempty"`

	// Policy-type-specific settings such as lease limits for lease policies.
	Definition interface{} `json:"definition,omitempty"`

	// definition legend
	DefinitionLegend map[string]DataElement `json:"definitionLegend,omitempty"`

	// The policy description.
	Description string `json:"description,omitempty"`

	// Defines enforcement type for policy. Default enforcement type is HARD.
	// Enum: [SOFT HARD]
	EnforcementType string `json:"enforcementType,omitempty"`

	// The policy ID.
	// Format: uuid
	ID strfmt.UUID `json:"id,omitempty"`

	// Most recent policy update timestamp.
	// Format: date-time
	LastUpdatedAt strfmt.DateTime `json:"lastUpdatedAt,omitempty"`

	// Most recent policy editor.
	LastUpdatedBy string `json:"lastUpdatedBy,omitempty"`

	// The policy name.
	Name string `json:"name,omitempty"`

	// The ID of the organization to which the policy belongs.
	OrgID string `json:"orgId,omitempty"`

	// For project-scoped policies, the ID of the project to which the policy belongs.
	ProjectID string `json:"projectId,omitempty"`

	// Project-based scope criteria to apply policy to multiple projects in the organization. Not allowed for project-scoped policies.
	ScopeCriteria *Criteria `json:"scopeCriteria,omitempty"`

	// statistics
	Statistics *PolicyStats `json:"statistics,omitempty"`

	// The policy type ID.
	// Required: true
	TypeID *string `json:"typeId"`
}


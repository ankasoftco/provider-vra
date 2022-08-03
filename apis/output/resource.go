package models 
type Resource struct {

	// Creation time
	// Format: date-time
	CreatedAt strfmt.DateTime `json:"createdAt,omitempty"`

	// Current ongoing request on the resource
	CurrentRequest *Request `json:"currentRequest,omitempty"`

	// Deployment to which resource belongs
	Deployment *DeploymentReference `json:"deployment,omitempty"`

	// Resource deployment id
	// Format: uuid
	DeploymentID strfmt.UUID `json:"deploymentId,omitempty"`

	// A description of the resource
	Description string `json:"description,omitempty"`

	// Expense associated with the deployment.
	// Read Only: true
	Expense *Expense `json:"expense,omitempty"`

	// Unique identifier of the resource
	// Format: uuid
	ID strfmt.UUID `json:"id,omitempty"`

	// Name of the resource
	// Required: true
	Name *string `json:"name"`

	// Resource org id
	OrgID string `json:"orgId,omitempty"`

	// Origin of the resource
	// Enum: [DISCOVERED ONBOARDED MIGRATED]
	Origin string `json:"origin,omitempty"`

	// Project to which resource's deployment belongs
	Project *ResourceReference `json:"project,omitempty"`

	// Resource project id
	ProjectID string `json:"projectId,omitempty"`

	// properties
	Properties interface{} `json:"properties,omitempty"`

	// The current sync status
	// Enum: [SUCCESS MISSING STALE]
	SyncStatus string `json:"syncStatus,omitempty"`

	// Type of the resource
	// Required: true
	Type *string `json:"type"`
}


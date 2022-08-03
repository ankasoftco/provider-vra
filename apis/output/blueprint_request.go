package models 
type BlueprintRequest struct {

	// Blueprint Id
	// Format: uuid
	BlueprintID strfmt.UUID `json:"blueprintId,omitempty"`

	// Blueprint version
	BlueprintVersion string `json:"blueprintVersion,omitempty"`

	// Cancel request time
	// Read Only: true
	// Format: date-time
	CancelRequestedAt strfmt.DateTime `json:"cancelRequestedAt,omitempty"`

	// Cancel requested by
	// Read Only: true
	CancelRequestedBy string `json:"cancelRequestedBy,omitempty"`

	// Blueprint YAML content
	Content string `json:"content,omitempty"`

	// Created time
	// Read Only: true
	// Format: date-time
	CreatedAt strfmt.DateTime `json:"createdAt,omitempty"`

	// Created by
	// Read Only: true
	CreatedBy string `json:"createdBy,omitempty"`

	// Existing deployment Id
	DeploymentID string `json:"deploymentId,omitempty"`

	// Name for the new deployment
	DeploymentName string `json:"deploymentName,omitempty"`

	// Description for the new request
	Description string `json:"description,omitempty"`

	// Destroy existing deployment
	Destroy bool `json:"destroy,omitempty"`

	// Failure message
	// Read Only: true
	FailureMessage string `json:"failureMessage,omitempty"`

	// Flow execution Id
	// Read Only: true
	FlowExecutionID string `json:"flowExecutionId,omitempty"`

	// Flow Id
	// Read Only: true
	FlowID string `json:"flowId,omitempty"`

	// Object ID
	// Read Only: true
	ID string `json:"id,omitempty"`

	// Ignore delete failures in blueprint request
	IgnoreDeleteFailures bool `json:"ignoreDeleteFailures,omitempty"`

	// Blueprint request inputs
	Inputs interface{} `json:"inputs,omitempty"`

	// Org ID
	// Read Only: true
	OrgID string `json:"orgId,omitempty"`

	// Plan only without affecting existing deployment
	Plan bool `json:"plan,omitempty"`

	// Project ID
	ProjectID string `json:"projectId,omitempty"`

	// Project Name
	// Read Only: true
	ProjectName string `json:"projectName,omitempty"`

	// Reason for requesting a blueprint
	Reason string `json:"reason,omitempty"`

	// Request tracker Id
	// Read Only: true
	RequestTrackerID string `json:"requestTrackerId,omitempty"`

	// Simulate blueprint request with providers
	Simulate bool `json:"simulate,omitempty"`

	// Status
	// Read Only: true
	// Enum: [CREATED STARTED FINISHED FAILED CANCELLED]
	Status string `json:"status,omitempty"`

	// Target resources
	// Read Only: true
	TargetResources []string `json:"targetResources"`

	// Updated time
	// Read Only: true
	// Format: date-time
	UpdatedAt strfmt.DateTime `json:"updatedAt,omitempty"`

	// Updated by
	// Read Only: true
	UpdatedBy string `json:"updatedBy,omitempty"`

	// Validation messages
	// Read Only: true
	ValidationMessages []*BlueprintValidationMessage `json:"validationMessages"`
}


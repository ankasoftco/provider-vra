package models 
type Request struct {

	// Identifier of the requested action
	ActionID string `json:"actionId,omitempty"`

	// Time at which the request was approved.
	// Format: date-time
	ApprovedAt strfmt.DateTime `json:"approvedAt,omitempty"`

	// Identifier of the requested blueprint in the form 'UUID:version'
	BlueprintID string `json:"blueprintId,omitempty"`

	// Indicates whether request can be canceled or not.
	Cancelable bool `json:"cancelable,omitempty"`

	// Identifier of the requested catalog item in the form 'UUID:version'
	CatalogItemID string `json:"catalogItemId,omitempty"`

	// Time at which the request completed.
	// Format: date-time
	CompletedAt strfmt.DateTime `json:"completedAt,omitempty"`

	// The number of tasks completed while fulfilling this request.
	// Required: true
	CompletedTasks *int32 `json:"completedTasks"`

	// Creation time (e.g. date format '2019-07-13T23:16:49.310Z').
	// Required: true
	// Format: date-time
	CreatedAt *strfmt.DateTime `json:"createdAt"`

	// Identifier of the requested deployment id to which the request applies to
	// Format: uuid
	DeploymentID strfmt.UUID `json:"deploymentId,omitempty"`

	// Longer user-friendly details of the request.
	Details string `json:"details,omitempty"`

	// Indicates whether request is in dismissed state.
	Dismissed bool `json:"dismissed,omitempty"`

	// Request identifier
	// Format: uuid
	ID strfmt.UUID `json:"id,omitempty"`

	// Time at which the request was initialized.
	// Format: date-time
	InitializedAt strfmt.DateTime `json:"initializedAt,omitempty"`

	// Request inputs
	Inputs interface{} `json:"inputs,omitempty"`

	// Short user-friendly label of the request (e.g. 'shuting down myVM')
	// Required: true
	Name *string `json:"name"`

	// Request outputs
	Outputs interface{} `json:"outputs,omitempty"`

	// User that initiated the request
	// Required: true
	RequestedBy *string `json:"requestedBy"`

	// Optional resource ids to which the request applies to
	ResourceIds []strfmt.UUID `json:"resourceIds"`

	// Request overall execution status.
	// Enum: [CREATED PENDING INITIALIZATION CHECKING_APPROVAL APPROVAL_PENDING INPROGRESS COMPLETION APPROVAL_REJECTED ABORTED SUCCESSFUL FAILED]
	Status string `json:"status,omitempty"`

	// The total number of tasks need to be completed to fulfil this request.
	// Required: true
	TotalTasks *int32 `json:"totalTasks"`

	// Last update time (e.g. date format '2019-07-13T23:16:49.310Z').
	// Format: date-time
	UpdatedAt strfmt.DateTime `json:"updatedAt,omitempty"`
}


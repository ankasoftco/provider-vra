package models 
type UserOpResource struct {

	// approver groups
	ApproverGroups []string `json:"approverGroups"`

	// approvers
	Approvers []string `json:"approvers"`

	// cancel previous pending user op
	CancelPreviousPendingUserOp bool `json:"cancelPreviousPendingUserOp,omitempty"`

	// change log
	ChangeLog string `json:"changeLog,omitempty"`

	// comments
	Comments string `json:"comments,omitempty"`

	// created at
	// Format: date-time
	CreatedAt strfmt.DateTime `json:"createdAt,omitempty"`

	// created by
	CreatedBy string `json:"createdBy,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// endpoint
	Endpoint string `json:"endpoint,omitempty"`

	// execution Id
	ExecutionID string `json:"executionId,omitempty"`

	// execution index
	ExecutionIndex int64 `json:"executionIndex,omitempty"`

	// expiration
	Expiration int32 `json:"expiration,omitempty"`

	// expiration in days
	ExpirationInDays int32 `json:"expirationInDays,omitempty"`

	// expiration in seconds
	ExpirationInSeconds int64 `json:"expirationInSeconds,omitempty"`

	// expiration unit
	// Enum: [MINUTES HOURS DAYS]
	ExpirationUnit string `json:"expirationUnit,omitempty"`

	// expires on in micros
	ExpiresOnInMicros int64 `json:"expiresOnInMicros,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// index
	Index string `json:"index,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// org Id
	OrgID string `json:"orgId,omitempty"`

	// parent Id
	ParentID string `json:"parentId,omitempty"`

	// pipeline
	Pipeline string `json:"pipeline,omitempty"`

	// pipeline Id
	PipelineID string `json:"pipelineId,omitempty"`

	// project Id
	ProjectID string `json:"projectId,omitempty"`

	// requested by
	RequestedBy string `json:"requestedBy,omitempty"`

	// requested on in micros
	RequestedOnInMicros int64 `json:"requestedOnInMicros,omitempty"`

	// responded by
	RespondedBy string `json:"respondedBy,omitempty"`

	// responded by email
	RespondedByEmail string `json:"respondedByEmail,omitempty"`

	// responded on in micros
	RespondedOnInMicros int64 `json:"respondedOnInMicros,omitempty"`

	// responder roles
	ResponderRoles []string `json:"responderRoles"`

	// responder token
	ResponderToken string `json:"responderToken,omitempty"`

	// resume count
	ResumeCount int32 `json:"resumeCount,omitempty"`

	// sendemail
	Sendemail bool `json:"sendemail,omitempty"`

	// stage key
	StageKey string `json:"stageKey,omitempty"`

	// status
	// Enum: [ACTIVE APPROVED CANCELED EXPIRED REJECTED]
	Status string `json:"status,omitempty"`

	// summary
	Summary string `json:"summary,omitempty"`

	// tags
	Tags []string `json:"tags"`

	// task execution Id
	TaskExecutionID string `json:"taskExecutionId,omitempty"`

	// task key
	TaskKey string `json:"taskKey,omitempty"`

	// updated at
	// Format: date-time
	UpdatedAt strfmt.DateTime `json:"updatedAt,omitempty"`

	// updated by
	UpdatedBy string `json:"updatedBy,omitempty"`

	// version
	Version string `json:"version,omitempty"`
}


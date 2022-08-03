package models 
type ResourceRequest struct {

	// action input properties
	ActionInputProperties interface{} `json:"actionInputProperties,omitempty"`

	// action name
	ActionName string `json:"actionName,omitempty"`

	// allocation resource
	AllocationResource interface{} `json:"allocationResource,omitempty"`

	// callback Url
	CallbackURL string `json:"callbackUrl,omitempty"`

	// deployment Id
	DeploymentID string `json:"deploymentId,omitempty"`

	// deployment name
	DeploymentName string `json:"deploymentName,omitempty"`

	// operation
	Operation string `json:"operation,omitempty"`

	// operation timeout seconds
	OperationTimeoutSeconds int64 `json:"operationTimeoutSeconds,omitempty"`

	// project Id
	ProjectID string `json:"projectId,omitempty"`

	// request Id
	RequestID string `json:"requestId,omitempty"`

	// resource link
	ResourceLink string `json:"resourceLink,omitempty"`

	// resource modified properties
	ResourceModifiedProperties interface{} `json:"resourceModifiedProperties,omitempty"`

	// resource name
	ResourceName string `json:"resourceName,omitempty"`

	// resource properties
	ResourceProperties interface{} `json:"resourceProperties,omitempty"`

	// resource request Id
	ResourceRequestID string `json:"resourceRequestId,omitempty"`

	// resource type
	ResourceType string `json:"resourceType,omitempty"`

	// tenant links
	TenantLinks []string `json:"tenantLinks"`
}


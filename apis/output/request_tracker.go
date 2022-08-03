package models 
type RequestTracker struct {

	// ID of the deployment, this request is connected to.
	// Example: 123e4567-e89b-12d3-a456-426614174000
	DeploymentID string `json:"deploymentId,omitempty"`

	// ID of this request.
	// Example: we655aew4e8e
	// Required: true
	ID *string `json:"id"`

	// Status message of the request.
	// Example: In Progress
	Message string `json:"message,omitempty"`

	// Name of the operation.
	// Example: Power-off
	Name string `json:"name,omitempty"`

	// Progress of the request as percentage.
	// Example: 90
	// Required: true
	Progress *int32 `json:"progress"`

	// Collection of resources.
	// Example: [\"/resources/i-0be3655a12fd28e8e\", \"/resources/i-0be234a15ft8det3\"]
	Resources []string `json:"resources"`

	// Self link of this request.
	// Example: /.../request-tracker/we655aew4e8e
	// Required: true
	SelfLink *string `json:"selfLink"`

	// Status of the request.
	// Example: FINISHED
	// Required: true
	// Enum: [FINISHED INPROGRESS FAILED]
	Status *string `json:"status"`
}


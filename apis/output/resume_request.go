package models 
type ResumeRequest struct {

	// action name
	ActionName string `json:"actionName,omitempty"`

	// id
	// Format: uuid
	ID strfmt.UUID `json:"id,omitempty"`

	// operation
	Operation string `json:"operation,omitempty"`

	// request Id
	RequestID string `json:"requestId,omitempty"`

	// resource type
	ResourceType string `json:"resourceType,omitempty"`
}


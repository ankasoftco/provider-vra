package models 
type BlueprintValidationRequest struct {

	// Blueprint Id
	// Format: uuid
	BlueprintID strfmt.UUID `json:"blueprintId,omitempty"`

	// Blueprint Version
	BlueprintVersion string `json:"blueprintVersion,omitempty"`

	// Blueprint YAML content
	Content string `json:"content,omitempty"`

	// Blueprint request inputs
	Inputs interface{} `json:"inputs,omitempty"`

	// Project Id
	ProjectID string `json:"projectId,omitempty"`
}


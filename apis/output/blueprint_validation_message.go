package models 
type BlueprintValidationMessage struct {

	// Validation message
	// Read Only: true
	Message string `json:"message,omitempty"`

	// Metadata
	// Read Only: true
	Metadata map[string]string `json:"metadata,omitempty"`

	// Validation path
	// Read Only: true
	Path string `json:"path,omitempty"`

	// Resource name
	// Read Only: true
	ResourceName string `json:"resourceName,omitempty"`

	// Message type
	// Read Only: true
	// Enum: [INFO WARNING ERROR]
	Type string `json:"type,omitempty"`
}


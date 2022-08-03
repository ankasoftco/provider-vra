package models 
type BlueprintValidationResponse struct {

	// Is blueprint valid
	// Read Only: true
	Valid *bool `json:"valid,omitempty"`

	// Validation messages
	// Read Only: true
	ValidationMessages []*BlueprintValidationMessage `json:"validationMessages"`
}


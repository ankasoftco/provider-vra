package models 
type GitWebhookFileFilter struct {

	// Identifies the type of File Filter, that can be REGEX or PLAIN.
	// Example: PLAIN
	// Enum: [REGEX PLAIN]
	Type string `json:"type,omitempty"`

	// The value can be either plain text or a regex.
	// Example: example.txt
	Value string `json:"value,omitempty"`
}


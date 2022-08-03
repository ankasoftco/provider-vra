package models 
type Event struct {

	// Longer user-friendly details of the event.
	Details string `json:"details,omitempty"`

	// Indicates whether the event has logs or not.
	HasLogs bool `json:"hasLogs,omitempty"`

	// Event identifier
	// Format: uuid
	ID strfmt.UUID `json:"id,omitempty"`

	// Short user-friendly label of the event (e.g. 'shutting down myVM')
	// Required: true
	Name *string `json:"name"`

	// Optional resource name to which the event applies to
	ResourceName string `json:"resourceName,omitempty"`

	// Optional resource type to which the event applies to
	ResourceType string `json:"resourceType,omitempty"`

	// Timestamp of the Event (e.g. date format '2019-07-13T23:16:49.310Z').
	// Required: true
	// Format: date-time
	Timestamp *strfmt.DateTime `json:"timestamp"`
}


package models 
type EventLog struct {

	// Indicates whether this is the last log entry.
	EOF bool `json:"eof,omitempty"`

	// Event Log identifier
	// Format: uuid
	ID strfmt.UUID `json:"id,omitempty"`

	// Message of the event log
	// Required: true
	Message *string `json:"message"`

	// Row number of the Event Log.
	Rownum int32 `json:"rownum,omitempty"`

	// Timestamp of the Event log (e.g. date format '2019-07-13T23:16:49.310Z').
	// Required: true
	// Format: date-time
	Timestamp *strfmt.DateTime `json:"timestamp"`
}


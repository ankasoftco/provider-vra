package models 
type TimeSeriesValue struct {

	// The timestamp of the metric.
	// Read Only: true
	// Format: date-time
	Timestamp strfmt.DateTime `json:"timestamp,omitempty"`

	// The expense value at the timestamp.
	// Read Only: true
	Value float64 `json:"value,omitempty"`
}


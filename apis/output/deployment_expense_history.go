package models 
type DeploymentExpenseHistory struct {

	// The currency code of the expense history.
	// Read Only: true
	Currency string `json:"currency,omitempty"`

	// The list of expense history metric.
	// Read Only: true
	Data []*TimeSeriesValue `json:"data"`

	// The requested interval type.
	// Read Only: true
	// Enum: [daily weekly monthly]
	Interval string `json:"interval,omitempty"`
}


package models 
type Expense struct {

	// Additional expense incurred for the resource.
	// Read Only: true
	AdditionalExpense float64 `json:"additionalExpense,omitempty"`

	// Expense sync message code if any.
	// Read Only: true
	Code string `json:"code,omitempty"`

	// Compute expense of the resource.
	// Read Only: true
	ComputeExpense float64 `json:"computeExpense,omitempty"`

	// Last expense sync time.
	// Read Only: true
	// Format: date-time
	LastUpdatedTime strfmt.DateTime `json:"lastUpdatedTime,omitempty"`

	// Expense sync message if any.
	// Read Only: true
	Message string `json:"message,omitempty"`

	// Network expense of the resource.
	// Read Only: true
	NetworkExpense float64 `json:"networkExpense,omitempty"`

	// Storage expense of the resource.
	// Read Only: true
	StorageExpense float64 `json:"storageExpense,omitempty"`

	// Total expense of the resource.
	// Read Only: true
	TotalExpense float64 `json:"totalExpense,omitempty"`

	// Monetary unit.
	// Read Only: true
	Unit string `json:"unit,omitempty"`
}


package models 
type MarketplaceFilter struct {

	// entries
	Entries *MarketplaceFilterEntries `json:"entries,omitempty"`

	// Filter Id
	ID string `json:"id,omitempty"`

	// Filter Display Name
	Name string `json:"name,omitempty"`

	// Filter Type
	// Enum: [COLLECTION DATE NUMBER]
	Type string `json:"type,omitempty"`

	// Filter Value Type
	// Enum: [RANGE SINGLE MULTIPLE]
	ValueType string `json:"valueType,omitempty"`
}


package models 
type CatalogItemResourceUpfrontPriceResponse struct {

	// Additional Price incurred for the catalog.
	// Read Only: true
	DailyAdditionalPrice float64 `json:"dailyAdditionalPrice,omitempty"`

	// Compute Price of the catalog resource.
	// Read Only: true
	DailyComputePrice float64 `json:"dailyComputePrice,omitempty"`

	// Network Price of the catalog resource.
	// Read Only: true
	DailyNetworkPrice float64 `json:"dailyNetworkPrice,omitempty"`

	// Storage Price of the catalog resource.
	// Read Only: true
	DailyStoragePrice float64 `json:"dailyStoragePrice,omitempty"`

	// Total Price of the catalog resource.
	// Read Only: true
	DailyTotalPrice float64 `json:"dailyTotalPrice,omitempty"`

	// resource name
	ResourceName string `json:"resourceName,omitempty"`

	// resource type
	ResourceType string `json:"resourceType,omitempty"`

	// Id
	// Read Only: true
	ResourceUpfrontPriceID string `json:"resourceUpfrontPriceId,omitempty"`

	// Upfront price sync status
	// Read Only: true
	// Enum: [SUCCESS ERROR DATA_NOT_AVAILABLE CURRENCY_NOT_SET]
	Status string `json:"status,omitempty"`

	// Upfront price status detail.
	// Read Only: true
	StatusDetails string `json:"statusDetails,omitempty"`

	// Monetary unit.
	// Read Only: true
	Unit string `json:"unit,omitempty"`
}


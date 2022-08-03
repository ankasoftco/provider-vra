package models 
type CatalogItemUpfrontPriceResponse struct {

	// Additional Price incurred for the catalog.
	// Read Only: true
	DailyAdditionalPrice float64 `json:"dailyAdditionalPrice,omitempty"`

	// Compute Price of the catalog.
	// Read Only: true
	DailyComputePrice float64 `json:"dailyComputePrice,omitempty"`

	// Network Price of the catalog.
	// Read Only: true
	DailyNetworkPrice float64 `json:"dailyNetworkPrice,omitempty"`

	// Storage Price of the catalog.
	// Read Only: true
	DailyStoragePrice float64 `json:"dailyStoragePrice,omitempty"`

	// Total Price of the catalog.
	// Read Only: true
	DailyTotalPrice float64 `json:"dailyTotalPrice,omitempty"`

	// resource price details
	ResourcePriceDetails []*CatalogItemResourceUpfrontPriceResponse `json:"resourcePriceDetails"`

	// Upfront price sync status
	// Read Only: true
	// Enum: [STARTED PREPARING_COST_ESTIMATION IN_PROGRESS SUCCESS ERROR DATA_NOT_AVAILABLE CURRENCY_NOT_SET PUBLIC_CLOUD_NOT_SUPPORTED]
	Status string `json:"status,omitempty"`

	// Upfront price status detail.
	// Read Only: true
	StatusDetails string `json:"statusDetails,omitempty"`

	// Monetary unit.
	// Read Only: true
	Unit string `json:"unit,omitempty"`

	// Id
	// Read Only: true
	UpfrontPriceID string `json:"upfrontPriceId,omitempty"`
}


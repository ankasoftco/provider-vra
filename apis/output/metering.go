package models 
type Metering struct {

	// base rate
	BaseRate float64 `json:"baseRate,omitempty"`

	// charge based on
	// Enum: [USAGE]
	ChargeBasedOn string `json:"chargeBasedOn,omitempty"`

	// charge on power state
	// Enum: [ALWAYS ONLY_WHEN_POWERED_ON POWERED_ON_AT_LEAST_ONCE]
	ChargeOnPowerState string `json:"chargeOnPowerState,omitempty"`

	// charge period
	// Enum: [HOURLY DAILY WEEKLY MONTHLY]
	ChargePeriod string `json:"chargePeriod,omitempty"`

	// fixed price
	FixedPrice float64 `json:"fixedPrice,omitempty"`

	// unit
	Unit string `json:"unit,omitempty"`
}


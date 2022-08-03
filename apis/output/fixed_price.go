package models 
type FixedPrice struct {

	// charge period
	// Enum: [HOURLY DAILY WEEKLY MONTHLY]
	ChargePeriod string `json:"chargePeriod,omitempty"`

	// rate
	Rate float64 `json:"rate,omitempty"`
}


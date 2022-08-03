package models 
type OneTimeMeteringItem struct {

	// item name
	ItemName string `json:"itemName,omitempty"`

	// one time metering
	OneTimeMetering *OneTimeMetering `json:"oneTimeMetering,omitempty"`
}


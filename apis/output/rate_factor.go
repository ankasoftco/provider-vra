package models 
type RateFactor struct {

	// context metering item
	ContextMeteringItem string `json:"contextMeteringItem,omitempty"`

	// rate factor
	RateFactor float64 `json:"rateFactor,omitempty"`
}


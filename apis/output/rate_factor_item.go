package models 
type RateFactorItem struct {

	// key
	Key string `json:"key,omitempty"`

	// rate factor
	RateFactor *RateFactor `json:"rateFactor,omitempty"`

	// value
	Value string `json:"value,omitempty"`
}


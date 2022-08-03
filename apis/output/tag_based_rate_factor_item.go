package models 
type TagBasedRateFactorItem struct {

	// item name
	ItemName string `json:"itemName,omitempty"`

	// rate factors
	RateFactors []*RateFactorItem `json:"rateFactors"`
}


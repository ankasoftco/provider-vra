package models 
type MeteringItem struct {

	// item name
	ItemName string `json:"itemName,omitempty"`

	// metering
	Metering *Metering `json:"metering,omitempty"`
}


package models 
type TagBasedOneTimeMeteringItem struct {

	// item name
	ItemName string `json:"itemName,omitempty"`

	// one time meterings
	OneTimeMeterings []*TagBasedOneTimeMetering `json:"oneTimeMeterings"`
}


package models 
type TagBasedMeteringItem struct {

	// item name
	ItemName string `json:"itemName,omitempty"`

	// tag based meterings
	TagBasedMeterings []*TagBasedMetering `json:"tagBasedMeterings"`
}


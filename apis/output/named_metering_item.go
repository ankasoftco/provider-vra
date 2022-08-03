package models 
type NamedMeteringItem struct {

	// item name
	ItemName string `json:"itemName,omitempty"`

	// named meterings
	NamedMeterings []*NamedMetering `json:"namedMeterings"`
}


package models 
type CloudAccountRegions struct {

	// A set of regions that can be enabled for this cloud account.
	ExternalRegions []*RegionSpecification `json:"externalRegions"`
}


package models 
type ImageProfileSpecification struct {

	// A human-friendly description.
	Description string `json:"description,omitempty"`

	// Image mapping defined for the corresponding region.
	// Example: { \"ubuntu\": { \"id\": \"9e49\", \"name\": \"ami-ubuntu-16.04-1.9.1-00-1516139717\"}, \"centos\": { \"id\": \"9e50\", \"name\": \"ami-centos-7-1.13.0-00-1543963388\"}}
	// Required: true
	ImageMapping map[string]FabricImageDescription `json:"imageMapping"`

	// A human-friendly name used as an identifier in APIs that support this option.
	// Required: true
	Name *string `json:"name"`

	// The id of the region for which this profile is created
	// Example: 9e49
	// Required: true
	RegionID *string `json:"regionId"`
}


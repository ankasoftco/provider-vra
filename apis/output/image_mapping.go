package models 
type ImageMapping struct {

	// HATEOAS of the entity
	// Required: true
	Links map[string]Href `json:"_links"`

	// The id of the region for which this mapping is defined.
	// Example: us-east-1
	ExternalRegionID string `json:"externalRegionId,omitempty"`

	// Image mapping defined for the particular region.
	// Example: { \"ubuntu\" : {\"externalRegionId\": \"us-east-1\", \"externalId\": \"ami-2c7b5656\", \"name\": \"ami-ubuntu-16.04\", \"description\": \"instance-store - HVM - supports - UTF-8 - tools\"} }
	// Required: true
	Mapping map[string]ImageMappingDescription `json:"mapping"`
}


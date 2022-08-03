package models 
type FlavorMapping struct {

	// HATEOAS of the entity
	// Required: true
	Links map[string]Href `json:"_links"`

	// The id of the region for which this mapping is defined.
	// Example: us-east-1
	ExternalRegionID string `json:"externalRegionId,omitempty"`

	// Flavors defined for the particular region. Keyed by global flavor key.
	// Example: { \"small\": { \"name\": \"t2.small\", \"cpuCount\": \"1\", \"MemoryInMB\": \"2048\", \"storageType\": \"EBS\", \"networkType\": \"Low to Moderate\"} }
	// Required: true
	Mapping map[string]FabricFlavor `json:"mapping"`
}


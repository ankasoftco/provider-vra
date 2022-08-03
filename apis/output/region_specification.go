package models 
type RegionSpecification struct {

	// Unique identifier of region on the provider side.
	// Example: us-west
	// Required: true
	ExternalRegionID *string `json:"externalRegionId"`

	// Name of region on the provider side. In vSphere, the name of the region is different from its id.
	// Example: us-west
	// Required: true
	Name *string `json:"name"`
}


package models 
type ZoneSpecification struct {

	// The ids of the compute resources that will be explicitly assigned to this zone
	// Example: [ab12c]
	ComputeIds []string `json:"computeIds"`

	// A list of key value pair of properties that will  be used
	// Example: {\"__isDefaultPlacementZone\": \"true\"}
	CustomProperties map[string]string `json:"customProperties,omitempty"`

	// A human-friendly description.
	Description string `json:"description,omitempty"`

	// The folder relative path to the datacenter where resources are deployed to. (only applicable for vSphere cloud zones)
	// Example: test-folder
	Folder string `json:"folder,omitempty"`

	// A human-friendly name used as an identifier in APIs that support this option.
	// Required: true
	Name *string `json:"name"`

	// Placement policy for the zone. One of DEFAULT, SPREAD or BINPACK.
	// Example: DEFAULT
	PlacementPolicy string `json:"placementPolicy,omitempty"`

	// The id of the region for which this profile is created
	// Example: 9e49
	// Required: true
	RegionID *string `json:"regionId"`

	// A set of tag keys and optional values that are effectively applied to all compute resources in this zone, but only in the context of this zone.
	// Example: [ { \"key\" : \"production\", \"value\": \" \" } ]
	Tags []*Tag `json:"tags"`

	// A set of tag keys and optional values that will be used
	// Example: [ { \"key\" : \"compliance\", \"value\": \"pci\" } ]
	TagsToMatch []*Tag `json:"tagsToMatch"`
}


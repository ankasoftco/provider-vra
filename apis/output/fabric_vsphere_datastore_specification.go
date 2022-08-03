package models 
type FabricVsphereDatastoreSpecification struct {

	// A set of tag keys and optional values that were set on this resource instance.
	// Example: [ { \"key\" : \"fast-network\", \"value\": \"true\" } ]
	Tags []*Tag `json:"tags"`
}


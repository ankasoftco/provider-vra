package models 
type FabricComputeSpecification struct {

	// A set of tag keys and optional values that were set on this resource instance.
	// Example: [ { \"key\" : \"?\", \"value\": \"Environment\" } ]
	Tags []*Tag `json:"tags"`
}


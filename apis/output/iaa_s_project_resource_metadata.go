package models 
type IaaSProjectResourceMetadata struct {

	// A list of keys and optional values to be applied to compute resources provisioned in a project
	// Example: [ { \"key\" : \"env\", \"value\": \"dev\" } ]
	Tags []*Tag `json:"tags"`
}


package models 
type UpdateExternalNetworkIPRangeSpecification struct {

	// A list of fabric network Ids that this IP range should be associated with.
	// Unique: true
	FabricNetworkIds []string `json:"fabricNetworkIds"`
}


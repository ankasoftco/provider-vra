package models 
type NodeSpec struct {

	// external Id
	ExternalID string `json:"externalId,omitempty"`

	// pod c ID r
	PodCIDR string `json:"podCIDR,omitempty"`

	// provider Id
	ProviderID string `json:"providerId,omitempty"`

	// unschedulable
	Unschedulable bool `json:"unschedulable,omitempty"`
}


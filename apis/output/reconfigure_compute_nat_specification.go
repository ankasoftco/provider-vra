package models 
type ReconfigureComputeNatSpecification struct {

	// List of NAT rules to be applied on this Compute Nat.
	// Required: true
	NatRules []*NatRule `json:"natRules"`
}


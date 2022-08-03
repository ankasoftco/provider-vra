package models 
type NodeStatus struct {

	// allocatable
	Allocatable interface{} `json:"allocatable,omitempty"`

	// capacity
	Capacity interface{} `json:"capacity,omitempty"`

	// phase
	Phase string `json:"phase,omitempty"`
}


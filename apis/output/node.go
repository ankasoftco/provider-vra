package models 
type Node struct {

	// api version
	APIVersion string `json:"apiVersion,omitempty"`

	// kind
	Kind string `json:"kind,omitempty"`

	// metadata
	Metadata *ObjectMeta `json:"metadata,omitempty"`

	// spec
	Spec *NodeSpec `json:"spec,omitempty"`

	// status
	Status *NodeStatus `json:"status,omitempty"`
}


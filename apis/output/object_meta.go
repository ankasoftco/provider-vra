package models 
type ObjectMeta struct {

	// annotations
	Annotations map[string]string `json:"annotations,omitempty"`

	// cluster name
	ClusterName string `json:"clusterName,omitempty"`

	// creation timestamp
	CreationTimestamp string `json:"creationTimestamp,omitempty"`

	// deletion timestamp
	DeletionTimestamp string `json:"deletionTimestamp,omitempty"`

	// generate name
	GenerateName string `json:"generateName,omitempty"`

	// labels
	Labels map[string]string `json:"labels,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// namespace
	Namespace string `json:"namespace,omitempty"`

	// self link
	SelfLink string `json:"selfLink,omitempty"`

	// uid
	UID string `json:"uid,omitempty"`
}


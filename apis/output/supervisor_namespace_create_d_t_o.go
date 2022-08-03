package models 
type SupervisorNamespaceCreateDTO struct {

	// cluster
	Cluster string `json:"cluster,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// endpoint Id
	EndpointID string `json:"endpointId,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// project Id
	ProjectID string `json:"projectId,omitempty"`

	// storage policies
	StoragePolicies []*StorageSpec `json:"storagePolicies"`
}


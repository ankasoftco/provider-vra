package models 
type SupervisorNamespaceUpdateDTO struct {

	// description
	Description string `json:"description,omitempty"`

	// resource quotas
	ResourceQuotas []*SupervisorNamespaceQuota `json:"resourceQuotas"`

	// storage policies
	StoragePolicies []*StorageSpec `json:"storagePolicies"`
}


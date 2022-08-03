package models 
type SupervisorNamespace struct {

	// address
	Address string `json:"address,omitempty"`

	// cluster
	Cluster string `json:"cluster,omitempty"`

	// cpu used
	CPUUsed int64 `json:"cpuUsed,omitempty"`

	// custom properties
	CustomProperties map[string]string `json:"customProperties,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// document self link
	DocumentSelfLink string `json:"documentSelfLink,omitempty"`

	// edit groups
	EditGroups string `json:"editGroups,omitempty"`

	// edit users
	EditUsers string `json:"editUsers,omitempty"`

	// endpoint link
	EndpointLink string `json:"endpointLink,omitempty"`

	// external link
	ExternalLink string `json:"externalLink,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// memory used
	MemoryUsed int64 `json:"memoryUsed,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// owner
	Owner string `json:"owner,omitempty"`

	// project Id
	ProjectID string `json:"projectId,omitempty"`

	// registered
	Registered bool `json:"registered,omitempty"`

	// self link Id
	SelfLinkID string `json:"selfLinkId,omitempty"`

	// status
	// Enum: [READY ALLOCATED ERROR REMOVING]
	Status string `json:"status,omitempty"`

	// status message
	StatusMessage string `json:"statusMessage,omitempty"`

	// storage policies
	StoragePolicies []*StorageSpec `json:"storagePolicies"`

	// storage used
	StorageUsed int64 `json:"storageUsed,omitempty"`

	// type
	Type string `json:"type,omitempty"`

	// updated millis
	UpdatedMillis int64 `json:"updatedMillis,omitempty"`

	// view groups
	ViewGroups string `json:"viewGroups,omitempty"`

	// view users
	ViewUsers string `json:"viewUsers,omitempty"`
}


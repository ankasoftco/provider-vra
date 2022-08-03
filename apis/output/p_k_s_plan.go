package models 
type PKSPlan struct {

	// description
	Description string `json:"description,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// master instances
	MasterInstances string `json:"masterInstances,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// worker instances
	WorkerInstances string `json:"workerInstances,omitempty"`
}


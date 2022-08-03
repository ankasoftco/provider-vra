package models 
type K8SCluster struct {

	// address
	Address string `json:"address,omitempty"`

	// auth credentials link
	AuthCredentialsLink string `json:"authCredentialsLink,omitempty"`

	// ca certificate
	CaCertificate string `json:"caCertificate,omitempty"`

	// cluster type
	ClusterType string `json:"clusterType,omitempty"`

	// content
	Content string `json:"content,omitempty"`

	// created millis
	CreatedMillis int64 `json:"createdMillis,omitempty"`

	// credentials
	Credentials *AuthCredentialsServiceState `json:"credentials,omitempty"`

	// custom properties
	CustomProperties interface{} `json:"customProperties,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// direct connection
	DirectConnection bool `json:"directConnection,omitempty"`

	// endpoint Id
	EndpointID string `json:"endpointId,omitempty"`

	// external link
	ExternalLink string `json:"externalLink,omitempty"`

	// global
	Global bool `json:"global,omitempty"`

	// id
	// Format: uuid
	ID strfmt.UUID `json:"id,omitempty"`

	// installer Id
	// Format: uuid
	InstallerID strfmt.UUID `json:"installerId,omitempty"`

	// kubernetes worker instances
	KubernetesWorkerInstances int32 `json:"kubernetesWorkerInstances,omitempty"`

	// last operation
	LastOperation string `json:"lastOperation,omitempty"`

	// last operation message
	LastOperationMessage string `json:"lastOperationMessage,omitempty"`

	// last operation status
	LastOperationStatus string `json:"lastOperationStatus,omitempty"`

	// last refreshed millis
	LastRefreshedMillis int64 `json:"lastRefreshedMillis,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// nodes
	Nodes []*K8SNode `json:"nodes"`

	// org Id
	OrgID string `json:"orgId,omitempty"`

	// owner
	Owner string `json:"owner,omitempty"`

	// project Id
	ProjectID string `json:"projectId,omitempty"`

	// remaining unreachable retry threshold
	RemainingUnreachableRetryThreshold int32 `json:"remainingUnreachableRetryThreshold,omitempty"`

	// shared
	Shared bool `json:"shared,omitempty"`

	// status
	Status string `json:"status,omitempty"`

	// updated millis
	UpdatedMillis int64 `json:"updatedMillis,omitempty"`

	// version
	Version string `json:"version,omitempty"`
}


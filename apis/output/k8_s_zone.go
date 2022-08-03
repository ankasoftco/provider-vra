package models 
type K8SZone struct {

	// clusters
	Clusters []*K8SZoneClusterAssignment `json:"clusters"`

	// created millis
	CreatedMillis int64 `json:"createdMillis,omitempty"`

	// custom properties
	CustomProperties interface{} `json:"customProperties,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// id
	// Format: uuid
	ID strfmt.UUID `json:"id,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// org Id
	OrgID string `json:"orgId,omitempty"`

	// projects
	Projects []*K8SZoneProjectAssignment `json:"projects"`

	// provider Id
	ProviderID string `json:"providerId,omitempty"`

	// provider type
	// Enum: [PKS_ENDPOINT EXTERNAL OPEN_SHIFT VSPHERE_NAMESPACES]
	ProviderType string `json:"providerType,omitempty"`

	// resources
	Resources []*K8SZoneResourceAssignment `json:"resources"`

	// supervisor clusters
	SupervisorClusters []*SupervisorClusterAssignment `json:"supervisorClusters"`

	// supervisor namespaces
	SupervisorNamespaces []*SupervisorNamespaceAssignment `json:"supervisorNamespaces"`

	// tag links
	TagLinks []string `json:"tagLinks"`

	// tags
	Tags []*TagState `json:"tags"`

	// updated millis
	UpdatedMillis int64 `json:"updatedMillis,omitempty"`
}


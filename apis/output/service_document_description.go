package models 
type ServiceDocumentDescription struct {

	// description
	Description string `json:"description,omitempty"`

	// document indexing options
	DocumentIndexingOptions []string `json:"documentIndexingOptions"`

	// name
	Name string `json:"name,omitempty"`

	// property descriptions
	PropertyDescriptions map[string]PropertyDescription `json:"propertyDescriptions,omitempty"`

	// serialized state size limit
	SerializedStateSizeLimit int32 `json:"serializedStateSizeLimit,omitempty"`

	// service capabilities
	ServiceCapabilities []string `json:"serviceCapabilities"`

	// service request routes
	ServiceRequestRoutes map[string][]Route `json:"serviceRequestRoutes,omitempty"`

	// user interface resource path
	UserInterfaceResourcePath string `json:"userInterfaceResourcePath,omitempty"`

	// version retention floor
	VersionRetentionFloor int64 `json:"versionRetentionFloor,omitempty"`

	// version retention limit
	VersionRetentionLimit int64 `json:"versionRetentionLimit,omitempty"`
}


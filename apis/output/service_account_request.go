package models 
type ServiceAccountRequest struct {

	// host name
	HostName string `json:"hostName,omitempty"`

	// resource Id
	ResourceID string `json:"resourceId,omitempty"`

	// resource type
	ResourceType string `json:"resourceType,omitempty"`
}


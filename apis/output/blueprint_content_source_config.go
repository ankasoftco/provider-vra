package models 
type BlueprintContentSourceConfig struct {

	// branch
	Branch string `json:"branch,omitempty"`

	// content type
	ContentType string `json:"contentType,omitempty"`

	// integration Id
	IntegrationID string `json:"integrationId,omitempty"`

	// path
	Path string `json:"path,omitempty"`

	// repository
	Repository string `json:"repository,omitempty"`
}


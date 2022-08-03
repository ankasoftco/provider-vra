package models 
type ContentAbout struct {

	// Latest API Version
	// Read Only: true
	LatestAPIVersion string `json:"latestApiVersion,omitempty"`

	// Supported APIs
	// Read Only: true
	SupportedApis []*SupportedAPI `json:"supportedApis"`
}


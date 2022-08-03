package models 
type BlueprintAbout struct {

	// Latest API Version
	// Read Only: true
	LatestAPIVersion string `json:"latestApiVersion,omitempty"`

	// Supported API's
	// Read Only: true
	SupportedApis []*SupportedAPI `json:"supportedApis"`
}


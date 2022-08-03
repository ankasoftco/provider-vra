package models 
type IaaSAbout struct {

	// The latest version of the API in yyyy-MM-dd format (UTC).
	// Required: true
	LatestAPIVersion *string `json:"latestApiVersion"`

	// A collection of all currently supported api versions.
	// Required: true
	SupportedApis []*APIDescription `json:"supportedApis"`
}


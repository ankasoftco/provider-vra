package models 
type CatalogItemRequest struct {

	// Deployment request count; defaults to 1 if not specified.
	// Maximum: 127
	// Minimum: -128
	BulkRequestCount *int32 `json:"bulkRequestCount,omitempty"`

	// Name of the requested deployment
	DeploymentName string `json:"deploymentName,omitempty"`

	// Input parameters for the request. These must be compliant with the schema of the corresponding catalog item
	Inputs interface{} `json:"inputs,omitempty"`

	// Project to be used for the request
	ProjectID string `json:"projectId,omitempty"`

	// Reason for request
	Reason string `json:"reason,omitempty"`

	// Version of the catalog item. e.g. v2.0
	Version string `json:"version,omitempty"`
}


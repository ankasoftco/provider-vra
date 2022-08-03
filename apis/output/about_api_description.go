package models 
type AboutAPIDescription struct {

	// The version of the API in yyyy-MM-dd format (UTC).
	// Required: true
	APIVersion *string `json:"apiVersion"`

	// The deprecation policy may contain information whether the api is in deprecated state and when it expires.
	DeprecationPolicy *AboutDeprecationPolicy `json:"deprecationPolicy,omitempty"`

	// The link to the documentation of this api version
	// Required: true
	DocumentationLink *string `json:"documentationLink"`
}

type AboutAPITESTDescription struct {

	// The version of the API in yyyy-MM-dd format (UTC).
	// Required: true
	APIVersion *string `json:"apiVersion"`

	// The deprecation policy may contain information whether the api is in deprecated state and when it expires.
	DeprecationPolicy *AboutDeprecationPolicy `json:"deprecationPolicy,omitempty"`

	// The link to the documentation of this api version
	// Required: true
	DocumentationLink *string `json:"documentationLink"`
}


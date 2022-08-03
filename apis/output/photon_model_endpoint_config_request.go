package models 
type PhotonModelEndpointConfigRequest struct {

	// associated endpoint links
	AssociatedEndpointLinks []string `json:"associatedEndpointLinks"`

	// auth credentials link
	AuthCredentialsLink string `json:"authCredentialsLink,omitempty"`

	// check for endpoint uniqueness
	CheckForEndpointUniqueness bool `json:"checkForEndpointUniqueness,omitempty"`

	// custom properties
	CustomProperties *CustomProperties `json:"customProperties,omitempty"`

	// endpoint properties
	EndpointProperties *EndpointProperties `json:"endpointProperties,omitempty"`

	// endpoint type
	EndpointType string `json:"endpointType,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// request type
	// Enum: [VALIDATE ENHANCE CHECK_IF_ACCOUNT_EXISTS VALIDATE_REMOVAL]
	RequestType string `json:"requestType,omitempty"`

	// resource link
	ResourceLink string `json:"resourceLink,omitempty"`

	// resource reference
	ResourceReference *URI `json:"resourceReference,omitempty"`

	// task reference
	TaskReference *URI `json:"taskReference,omitempty"`
}


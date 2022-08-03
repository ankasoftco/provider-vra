package models 
type FormDefinition struct {

	// form URI
	FormURI string `json:"formURI,omitempty"`

	// track progress in modal
	TrackProgressInModal bool `json:"trackProgressInModal,omitempty"`
}


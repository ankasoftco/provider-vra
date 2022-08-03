package models 
type VcfServiceCredentialCreationResponse struct {

	// credentials
	Credentials []*VcfServiceAccount `json:"credentials"`

	// domain Id
	DomainID string `json:"domainId,omitempty"`

	// sddc integration Id
	SddcIntegrationID string `json:"sddcIntegrationId,omitempty"`

	// status
	// Enum: [SUCCESSFUL IN_PROGRESS FAILED]
	Status string `json:"status,omitempty"`

	// task Id
	TaskID string `json:"taskId,omitempty"`
}


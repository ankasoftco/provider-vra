package models 
type VcfServiceCredential struct {

	// auth credential link
	AuthCredentialLink string `json:"authCredentialLink,omitempty"`

	// credential Id
	CredentialID string `json:"credentialId,omitempty"`

	// domain Id
	DomainID string `json:"domainId,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// integration Id
	IntegrationID string `json:"integrationId,omitempty"`

	// resource Id
	ResourceID string `json:"resourceId,omitempty"`

	// resource type
	ResourceType string `json:"resourceType,omitempty"`

	// role name
	RoleName string `json:"roleName,omitempty"`

	// status
	Status string `json:"status,omitempty"`

	// task Id
	TaskID string `json:"taskId,omitempty"`

	// user name
	UserName string `json:"userName,omitempty"`

	// vcf cloud account Id
	VcfCloudAccountID string `json:"vcfCloudAccountId,omitempty"`
}


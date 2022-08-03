package models 
type CustomProperties struct {

	// endpoint Id
	EndpointID string `json:"endpointId,omitempty"`

	// is external
	IsExternal bool `json:"isExternal,omitempty"`

	// service account Id
	ServiceAccountID string `json:"serviceAccountId,omitempty"`

	// service account name
	ServiceAccountName string `json:"serviceAccountName,omitempty"`

	// vcf domain Id
	VcfDomainID string `json:"vcfDomainId,omitempty"`

	// vcf domain name
	VcfDomainName string `json:"vcfDomainName,omitempty"`
}


package models 
type UpdateCloudAccountVcfSpecification struct {

	// Accept self signed certificate when connecting to vSphere and NSX-T
	// Example: false
	AcceptSelfSignedCertificate bool `json:"acceptSelfSignedCertificate,omitempty"`

	// Certificate for a cloud account.
	// Example: {\"certificate\": \"-----BEGIN CERTIFICATE-----\\nMIIDHjCCAoegAwIBAgIBATANBgkqhkiG9w0BAQsFADCBpjEUMBIGA1UEChMLVk13\\nYXJlIEluYAAc1pw18GT3iAqQRPx0PrjzJhgjIJMla\\n/1Kg4byY4FPSacNiRgY/FG2bPCqZk1yRfzmkFYCW/vU+Dg==\\n-----END CERTIFICATE-----\\n-\"}
	CertificateInfo *CertificateInfoSpecification `json:"certificateInfo,omitempty"`

	// Create default cloud zones for the enabled regions.
	// Example: true
	CreateDefaultZones bool `json:"createDefaultZones,omitempty"`

	// Identifier of a data collector vm deployed in the on premise infrastructure. Refer to the data-collector API to create or list data collectors.
	// Note: Data collector endpoints are not available in vRA on-prem release.
	// Example: 23959a1e-18bc-4f0c-ac49-b5aeb4b6eef4
	DcID string `json:"dcId,omitempty"`

	// A human-friendly description.
	Description string `json:"description,omitempty"`

	// A human-friendly name used as an identifier in APIs that support this option.
	Name string `json:"name,omitempty"`

	// NSX Certificate
	NsxCertificate string `json:"nsxCertificate,omitempty"`

	// Host name for the NSX endpoint from the specified workload domain.
	// Example: nsxt.mycompany.com
	// Required: true
	NsxHostName *string `json:"nsxHostName"`

	// Password for the user used to authenticate with the NSX-T manager in VCF cloud account
	// Example: cndhjslacd90ascdbasyoucbdh
	// Required: true
	NsxPassword *string `json:"nsxPassword"`

	// User name for the NSX manager in the specified workload domain.
	// Example: administrator@mycompany.com
	// Required: true
	NsxUsername *string `json:"nsxUsername"`

	// A set of regions to enable provisioning on.Refer to /iaas/api/cloud-accounts/region-enumeration.
	// Example: [{ \"name\": \"us-east-1\",\"externalRegionId\": \"us-east-1\"}]
	Regions []*RegionSpecification `json:"regions"`

	// SDDC manager integration id
	SddcManagerID string `json:"sddcManagerId,omitempty"`

	// A set of tag keys and optional values to set on the Cloud Account.Cloud account capability tags may enable different features.
	// Example: [ { \"key\" : \"env\", \"value\": \"dev\" } ]
	Tags []*Tag `json:"tags"`

	// vCenter Certificate
	VcenterCertificate string `json:"vcenterCertificate,omitempty"`

	// Host name for the vSphere from the specified workload domain.
	// Example: vc.mycompany.com
	// Required: true
	VcenterHostName *string `json:"vcenterHostName"`

	// Password for the user used to authenticate with the vCenter in VCF cloud account
	// Example: cndhjslacd90ascdbasyoucbdh
	// Required: true
	VcenterPassword *string `json:"vcenterPassword"`

	// vCenter user name for the specified workload domain.The specified user requires CloudAdmin credentials. The user does not require CloudGlobalAdmin credentials.
	// Example: administrator@mycompany.com
	// Required: true
	VcenterUsername *string `json:"vcenterUsername"`

	// Id of the workload domain to add as VCF cloud account.
	// Required: true
	WorkloadDomainID *string `json:"workloadDomainId"`

	// Name of the workload domain to add as VCF cloud account.
	// Example: Management
	// Required: true
	WorkloadDomainName *string `json:"workloadDomainName"`
}


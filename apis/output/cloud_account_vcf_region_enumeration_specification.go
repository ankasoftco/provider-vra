package models 
type CloudAccountVcfRegionEnumerationSpecification struct {

	// Accept self signed certificate when connecting to vSphere and NSX-T
	// Example: false
	AcceptSelfSignedCertificate bool `json:"acceptSelfSignedCertificate,omitempty"`

	// Certificate for a cloud account.
	// Example: {\"certificate\": \"-----BEGIN CERTIFICATE-----\\nMIIDHjCCAoegAwIBAgIBATANBgkqhkiG9w0BAQsFADCBpjEUMBIGA1UEChMLVk13\\nYXJlIEluYAAc1pw18GT3iAqQRPx0PrjzJhgjIJMla\\n/1Kg4byY4FPSacNiRgY/FG2bPCqZk1yRfzmkFYCW/vU+Dg==\\n-----END CERTIFICATE-----\\n-\"}
	CertificateInfo *CertificateInfoSpecification `json:"certificateInfo,omitempty"`

	// Existing cloud account id. Either provide existing cloud account Id, or workloadDomainId, workloadDomainName, vcenterHostName, vcenterUsername, vcenterPassword, nsxHostName, nsxUsername and nsxPassword.
	// Example: b8b7a918-342e-4a53-a3b0-b935da0fe601
	CloudAccountID string `json:"cloudAccountId,omitempty"`

	// Identifier of a data collector vm deployed in the on premise infrastructure. Refer to the data-collector API to create or list data collectors.
	// Note: Data collector endpoints are not available in vRA on-prem release.
	// Example: 23959a1e-18bc-4f0c-ac49-b5aeb4b6eef4
	DcID string `json:"dcId,omitempty"`

	// NSX Certificate
	NsxCertificate string `json:"nsxCertificate,omitempty"`

	// Host name for the NSX endpoint from the specified workload domain. Either provide nsxHostName or provide a cloudAccountId of an existing account.
	// Example: nsxt.mycompany.com
	NsxHostName string `json:"nsxHostName,omitempty"`

	// Password for the user used to authenticate with the NSX-T manager in VCF cloud account. Either provide nsxPassword or provide a cloudAccountId of an existing account.
	// Example: cndhjslacd90ascdbasyoucbdh
	NsxPassword string `json:"nsxPassword,omitempty"`

	// User name for the NSX manager in the specified workload domain. Either provide nsxUsername or provide a cloudAccountId of an existing account.
	// Example: administrator@mycompany.com
	NsxUsername string `json:"nsxUsername,omitempty"`

	// SDDC manager integration id. Either provide sddcManagerId or provide a cloudAccountId of an existing account.
	SddcManagerID string `json:"sddcManagerId,omitempty"`

	// vCenter Certificate
	VcenterCertificate string `json:"vcenterCertificate,omitempty"`

	// Host name for the vSphere from the specified workload domain. Either provide vcenterHostName or provide a cloudAccountId of an existing account.
	// Example: vc.mycompany.com
	VcenterHostName string `json:"vcenterHostName,omitempty"`

	// Password for the user used to authenticate with the vCenter in VCF cloud account.  Either provide vcenterPassword or provide a cloudAccountId of an existing account.
	// Example: cndhjslacd90ascdbasyoucbdh
	VcenterPassword string `json:"vcenterPassword,omitempty"`

	// vCenter user name for the specified workload domain.The specified user requires CloudAdmin credentials. The user does not require CloudGlobalAdmin credentials.
	// Example: administrator@mycompany.com
	VcenterUsername string `json:"vcenterUsername,omitempty"`

	// Id of the workload domain to add as VCF cloud account. Either provide workloadDomainId or provide a cloudAccountId of an existing account.
	WorkloadDomainID string `json:"workloadDomainId,omitempty"`

	// Name of the workload domain to add as VCF cloud account. Either provide workloadDomainName or provide a cloudAccountId of an existing account.
	// Example: Management
	WorkloadDomainName string `json:"workloadDomainName,omitempty"`
}


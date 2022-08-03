package models 
type UpdateCloudAccountVmcSpecification struct {

	// Accept self signed certificate when connecting to vSphere
	// Example: false
	AcceptSelfSignedCertificate bool `json:"acceptSelfSignedCertificate,omitempty"`

	// VMC API access key. Optional when updating.
	APIKey string `json:"apiKey,omitempty"`

	// Certificate for a cloud account.
	// Example: {\"certificate\": \"-----BEGIN CERTIFICATE-----\\nMIIDHjCCAoegAwIBAgIBATANBgkqhkiG9w0BAQsFADCBpjEUMBIGA1UEChMLVk13\\nYXJlIEluYAAc1pw18GT3iAqQRPx0PrjzJhgjIJMla\\n/1Kg4byY4FPSacNiRgY/FG2bPCqZk1yRfzmkFYCW/vU+Dg==\\n-----END CERTIFICATE-----\\n-\"}
	CertificateInfo *CertificateInfoSpecification `json:"certificateInfo,omitempty"`

	// Create default cloud zones for the enabled regions.
	// Example: true
	CreateDefaultZones bool `json:"createDefaultZones,omitempty"`

	// Identifier of a data collector vm deployed in the on premise infrastructure. Refer to the data-collector API to create or list data collectors
	// Example: 23959a1e-18bc-4f0c-ac49-b5aeb4b6eef4
	// Required: true
	DcID *string `json:"dcId"`

	// A human-friendly description.
	Description string `json:"description,omitempty"`

	// Enter the IP address or FQDN of the vCenter Server in the specified SDDC. The cloud proxy belongs on this vCenter.
	// Example: vc1.vmware.com
	// Required: true
	HostName *string `json:"hostName"`

	// A human-friendly name used as an identifier in APIs that support this option.
	Name string `json:"name,omitempty"`

	// The IP address of the NSX Manager server in the specified SDDC / FQDN.
	// Example: nsxManager.sddc-52-12-8-145.vmwaretest.com
	// Required: true
	NsxHostName *string `json:"nsxHostName"`

	// Password for the user used to authenticate with the cloud Account
	// Example: cndhjslacd90ascdbasyoucbdh
	// Required: true
	Password *string `json:"password"`

	// A set of regions to enable provisioning on.Refer to /iaas/api/cloud-accounts/region-enumeration.
	// Example: [{ \"name\": \"Datacenter:datacenter-3\",\"externalRegionId\": \"Datacenter:datacenter-3\"}]
	Regions []*RegionSpecification `json:"regions"`

	// Identifier of the on-premise SDDC to be used by this cloud account. Note that NSX-V SDDCs are not supported.
	// Example: CMBU-PRD-NSXT-M8GA-090319
	// Required: true
	SddcID *string `json:"sddcId"`

	// A set of tag keys and optional values to set on the Cloud Account.Cloud account capability tags may enable different features.
	// Example: [ { \"key\" : \"env\", \"value\": \"dev\" } ]
	Tags []*Tag `json:"tags"`

	// vCenter user name for the specified SDDC.The specified user requires CloudAdmin credentials. The user does not require CloudGlobalAdmin credentials.
	// Example: administrator@mycompany.com
	// Required: true
	Username *string `json:"username"`
}


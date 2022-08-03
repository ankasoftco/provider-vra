package models 
type CloudAccountVsphereSpecification struct {

	// Accept self signed certificate when connecting to vSphere
	// Example: false
	AcceptSelfSignedCertificate bool `json:"acceptSelfSignedCertificate,omitempty"`

	// NSX-V or NSX-T account to associate with this vSphere cloud account. vSphere cloud account can be a single NSX-V cloud account or a single NSX-T cloud account.
	// Example: [ \"42f3e0d199d134755684cd935435a\" ]
	AssociatedCloudAccountIds []string `json:"associatedCloudAccountIds"`

	// Certificate for a cloud account.
	// Example: {\"certificate\": \"-----BEGIN CERTIFICATE-----\\nMIIDHjCCAoegAwIBAgIBATANBgkqhkiG9w0BAQsFADCBpjEUMBIGA1UEChMLVk13\\nYXJlIEluYAAc1pw18GT3iAqQRPx0PrjzJhgjIJMla\\n/1Kg4byY4FPSacNiRgY/FG2bPCqZk1yRfzmkFYCW/vU+Dg==\\n-----END CERTIFICATE-----\\n-\"}
	CertificateInfo *CertificateInfoSpecification `json:"certificateInfo,omitempty"`

	// Create default cloud zones for the enabled regions.
	// Example: true
	CreateDefaultZones bool `json:"createDefaultZones,omitempty"`

	// Identifier of a data collector vm deployed in the on premise infrastructure. Refer to the data-collector API to create or list data collectors.
	// Note: Data collector endpoints are not available in vRA on-prem release.
	// Example: 23959a1e-18bc-4f0c-ac49-b5aeb4b6eef4
	Dcid string `json:"dcid,omitempty"`

	// A human-friendly description.
	Description string `json:"description,omitempty"`

	// Host name for the vSphere endpoint
	// Example: vc.mycompany.com
	// Required: true
	HostName *string `json:"hostName"`

	// A human-friendly name used as an identifier in APIs that support this option.
	// Required: true
	Name *string `json:"name"`

	// Password for the user used to authenticate with the cloud Account
	// Example: cndhjslacd90ascdbasyoucbdh
	// Required: true
	Password *string `json:"password"`

	// A set of regions to enable provisioning on.Refer to /iaas/api/cloud-accounts/region-enumeration.
	// Example: [{ \"name\": \"Datacenter:datacenter-3\",\"externalRegionId\": \"Datacenter:datacenter-3\"}]
	// Required: true
	Regions []*RegionSpecification `json:"regions"`

	// A set of tag keys and optional values to set on the Cloud Account
	// Example: [ { \"key\" : \"env\", \"value\": \"dev\" } ]
	Tags []*Tag `json:"tags"`

	// Username to authenticate with the cloud account
	// Example: administrator@mycompany.com
	// Required: true
	Username *string `json:"username"`
}


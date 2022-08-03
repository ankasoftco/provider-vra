package models 
type CloudAccountNsxTSpecification struct {

	// Accept self signed certificate when connecting.
	// Example: false
	AcceptSelfSignedCertificate bool `json:"acceptSelfSignedCertificate,omitempty"`

	// vSphere cloud accounts associated with this NSX-T cloud account.
	// Example: [ \"42f3e0d199d134755684cd935435a\" ]
	AssociatedCloudAccountIds []string `json:"associatedCloudAccountIds"`

	// Certificate for a cloud account.
	// Example: {\"certificate\": \"-----BEGIN CERTIFICATE-----\\nMIIDHjCCAoegAwIBAgIBATANBgkqhkiG9w0BAQsFADCBpjEUMBIGA1UEChMLVk13\\nYXJlIEluYAAc1pw18GT3iAqQRPx0PrjzJhgjIJMla\\n/1Kg4byY4FPSacNiRgY/FG2bPCqZk1yRfzmkFYCW/vU+Dg==\\n-----END CERTIFICATE-----\\n-\"}
	CertificateInfo *CertificateInfoSpecification `json:"certificateInfo,omitempty"`

	// Identifier of a data collector vm deployed in the on premise infrastructure. Refer to the data-collector API to create or list data collectors.
	// Note: Data collector endpoints are not available in vRA on-prem release and hence the data collector Id is optional for vRA on-prem.
	// Example: 23959a1e-18bc-4f0c-ac49-b5aeb4b6eef4
	// Required: true
	Dcid *string `json:"dcid"`

	// A human-friendly description.
	Description string `json:"description,omitempty"`

	// Host name for the NSX-T endpoint
	// Example: nsxt.mycompany.com
	// Required: true
	HostName *string `json:"hostName"`

	// Indicates whether this is an NSX-T Global Manager cloud account. NSX-T Global Manager can only be associated with NSX-T cloud accounts. Default value: false.
	// Example: true
	IsGlobalManager bool `json:"isGlobalManager,omitempty"`

	// Create NSX-T cloud account in Manager (legacy) mode. When set to true, NSX-T cloud account in created in Manager mode. Mode cannot be changed after cloud account is created. Default value is false.
	// Example: true
	ManagerMode bool `json:"managerMode,omitempty"`

	// A human-friendly name used as an identifier in APIs that support this option.
	// Required: true
	Name *string `json:"name"`

	// Password for the user used to authenticate with the cloud Account
	// Example: cndhjslacd90ascdbasyoucbdh
	// Required: true
	Password *string `json:"password"`

	// A set of tag keys and optional values to set on the Cloud Account
	// Example: [ { \"key\" : \"env\", \"value\": \"dev\" } ]
	Tags []*Tag `json:"tags"`

	// Username to authenticate with the cloud account
	// Example: administrator@mycompany.com
	// Required: true
	Username *string `json:"username"`
}


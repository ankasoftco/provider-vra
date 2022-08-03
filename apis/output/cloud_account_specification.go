package models 
type CloudAccountSpecification struct {

	// Cloud accounts to associate with this cloud account
	// Example: [ \"42f3e0d199d134755684cd935435a\" ]
	AssociatedCloudAccountIds []string `json:"associatedCloudAccountIds"`

	// Certificate for a cloud account.
	// Example: {\"certificate\": \"-----BEGIN CERTIFICATE-----\\nMIIDHjCCAoegAwIBAgIBATANBgkqhkiG9w0BAQsFADCBpjEUMBIGA1UEChMLVk13\\nYXJlIEluYAAc1pw18GT3iAqQRPx0PrjzJhgjIJMla\\n/1Kg4byY4FPSacNiRgY/FG2bPCqZk1yRfzmkFYCW/vU+Dg==\\n-----END CERTIFICATE-----\\n-\"}
	CertificateInfo *CertificateInfoSpecification `json:"certificateInfo,omitempty"`

	// Cloud Account specific properties supplied in as name value pairs
	// Example: {\"supportPublicImages\": \"true\", \"acceptSelfSignedCertificate\": \"true\" }
	// Required: true
	CloudAccountProperties map[string]string `json:"cloudAccountProperties"`

	// Cloud account type
	// Example: vsphere, aws, azure, nsxv, nsxt
	// Required: true
	CloudAccountType *string `json:"cloudAccountType"`

	// Create default cloud zones for the enabled regions.
	// Example: true
	CreateDefaultZones bool `json:"createDefaultZones,omitempty"`

	// Additional custom properties that may be used to extend the Cloud Account.
	// Example: { \"sampleadapterProjectId\" : \"projectId\" }
	CustomProperties map[string]string `json:"customProperties,omitempty"`

	// A human-friendly description.
	Description string `json:"description,omitempty"`

	// A human-friendly name used as an identifier in APIs that support this option.
	// Required: true
	Name *string `json:"name"`

	// Secret access key or password to be used to authenticate with the cloud account
	// Example: gfsScK345sGGaVdds222dasdfDDSSasdfdsa34fS
	// Required: true
	PrivateKey *string `json:"privateKey"`

	// Access key id or username to be used to authenticate with the cloud account
	// Example: ACDC55DB4MFH6ADG75KK
	// Required: true
	PrivateKeyID *string `json:"privateKeyId"`

	// A set of regions to enable provisioning on.Refer to /iaas/api/cloud-accounts/region-enumeration.
	// 'regionInfos' is a required parameter for AWS, AZURE, GCP, VSPHERE, VMC, VCF cloud account types.
	// Example: [{ \"name\": \"East Asia\",\"externalRegionId\": \"eastasia\"}]
	// Required: true
	Regions []*RegionSpecification `json:"regions"`

	// A set of tag keys and optional values to set on the Cloud Account
	// Example: [ { \"key\" : \"env\", \"value\": \"dev\" } ]
	Tags []*Tag `json:"tags"`
}


package models 
type CloudAccountRegionEnumerationSpecification struct {

	// Certificate for a cloud account.
	// Example: {\"certificate\": \"-----BEGIN CERTIFICATE-----\\nMIIDHjCCAoegAwIBAgIBATANBgkqhkiG9w0BAQsFADCBpjEUMBIGA1UEChMLVk13\\nYXJlIEluYAAc1pw18GT3iAqQRPx0PrjzJhgjIJMla\\n/1Kg4byY4FPSacNiRgY/FG2bPCqZk1yRfzmkFYCW/vU+Dg==\\n-----END CERTIFICATE-----\\n-\"}
	CertificateInfo *CertificateInfoSpecification `json:"certificateInfo,omitempty"`

	// Existing cloud account id. Either provide existing cloud account Id, or privateKeyId/privateKey credentials pair.
	// Example: b8b7a918-342e-4a53-a3b0-b935da0fe601
	CloudAccountID string `json:"cloudAccountId,omitempty"`

	// Cloud Account specific properties supplied in as name value pairs
	// Example: {\"supportPublicImages\": \"true\", \"acceptSelfSignedCertificate\": \"true\" }
	// Required: true
	CloudAccountProperties map[string]string `json:"cloudAccountProperties"`

	// Cloud account type
	// Example: vsphere, aws, azure, nsxv, nsxt
	CloudAccountType string `json:"cloudAccountType,omitempty"`

	// Secret access key or password to be used to authenticate with the cloud account. Either provide privateKey or provide a cloudAccountId of an existing account.
	// Example: gfsScK345sGGaVdds222dasdfDDSSasdfdsa34fS
	PrivateKey string `json:"privateKey,omitempty"`

	// Access key id or username to be used to authenticate with the cloud account. Either provide privateKeyId or provide a cloudAccountId of an existing account.
	// Example: ACDC55DB4MFH6ADG75KK
	PrivateKeyID string `json:"privateKeyId,omitempty"`
}


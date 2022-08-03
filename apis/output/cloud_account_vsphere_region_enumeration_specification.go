package models 
type CloudAccountVsphereRegionEnumerationSpecification struct {

	// Accept self signed certificate when connecting to vSphere
	// Example: false
	AcceptSelfSignedCertificate bool `json:"acceptSelfSignedCertificate,omitempty"`

	// Certificate for a cloud account.
	// Example: {\"certificate\": \"-----BEGIN CERTIFICATE-----\\nMIIDHjCCAoegAwIBAgIBATANBgkqhkiG9w0BAQsFADCBpjEUMBIGA1UEChMLVk13\\nYXJlIEluYAAc1pw18GT3iAqQRPx0PrjzJhgjIJMla\\n/1Kg4byY4FPSacNiRgY/FG2bPCqZk1yRfzmkFYCW/vU+Dg==\\n-----END CERTIFICATE-----\\n-\"}
	CertificateInfo *CertificateInfoSpecification `json:"certificateInfo,omitempty"`

	// Existing cloud account id. Either provide existing cloud account Id, or hostName, username, password.
	// Example: b8b7a918-342e-4a53-a3b0-b935da0fe601
	CloudAccountID string `json:"cloudAccountId,omitempty"`

	// Identifier of a data collector vm deployed in the on premise infrastructure. Refer to the data-collector API to create or list data collectors.
	// Note: Data collector endpoints are not available in vRA on-prem release.
	// Example: 23959a1e-18bc-4f0c-ac49-b5aeb4b6eef4
	Dcid string `json:"dcid,omitempty"`

	// Host name for the vSphere endpoint. Either provide hostName or provide a cloudAccountId of an existing account.
	// Example: vc.mycompany.com
	HostName string `json:"hostName,omitempty"`

	// Password for the user used to authenticate with the cloud Account. Either provide password or provide a cloudAccountId of an existing account.
	// Example: cndhjslacd90ascdbasyoucbdh
	Password string `json:"password,omitempty"`

	// Username to authenticate with the cloud account. Either provide username or provide a cloudAccountId of an existing account.
	// Example: administrator@mycompany.com
	Username string `json:"username,omitempty"`
}


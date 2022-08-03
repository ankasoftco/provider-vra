package models 
type UpdateIntegrationSpecification struct {

	// Cloud accounts to associate with this integration
	// Example: [ \"42f3e0d199d134755684cd935435a\" ]
	AssociatedCloudAccountIds []string `json:"associatedCloudAccountIds"`

	// Certificate for an integration.
	// Example: {\"certificate\": \"-----BEGIN CERTIFICATE-----\\nMIIDHjCCAoegAwIBAgIBATANBgkqhkiG9w0BAQsFADCBpjEUMBIGA1UEChMLVk13\\nYXJlIEluYAAc1pw18GT3iAqQRPx0PrjzJhgjIJMla\\n/1Kg4byY4FPSacNiRgY/FG2bPCqZk1yRfzmkFYCW/vU+Dg==\\n-----END CERTIFICATE-----\\n-\"}
	CertificateInfo *CertificateInfoSpecification `json:"certificateInfo,omitempty"`

	// Additional custom properties that may be used to extend the Integration.
	// Example: { \"sampleadapterProjectId\" : \"projectId\" }
	CustomProperties map[string]string `json:"customProperties,omitempty"`

	// A human-friendly description.
	Description string `json:"description,omitempty"`

	// Integration specific properties supplied in as name value pairs
	// Example: { \"providerId\" : \"providerID\" }
	IntegrationProperties map[string]string `json:"integrationProperties,omitempty"`

	// A human-friendly name used as an identifier in APIs that support this option.
	Name string `json:"name,omitempty"`

	// Secret access key or password to be used to authenticate with the integration
	// Example: gfsScK345sGGaVdds222dasdfDDSSasdfdsa34fS
	PrivateKey string `json:"privateKey,omitempty"`

	// Access key id or username to be used to authenticate with the integration
	// Example: ACDC55DB4MFH6ADG75KK
	PrivateKeyID string `json:"privateKeyId,omitempty"`

	// A set of tag keys and optional values to set on the Integration
	// Example: [ { \"key\" : \"env\", \"value\": \"dev\" } ]
	Tags []*Tag `json:"tags"`
}


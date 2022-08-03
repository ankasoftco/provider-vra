package models 
type CertificateInfo struct {

	// The certificate in string format.
	// Example: PEM for X509Certificate
	// Required: true
	Certificate *string `json:"certificate"`

	// Details about the certificate.
	// Example: UNTRUSTED_CERTIFICATE
	// Enum: [UNTRUSTED_CERTIFICATE EXPIRED_CERTIFICATE NOT_YET_VALID_CERTIFICATE KEYSTORE_TAMPERED_OR_PASSWORD_INCORRECT]
	CertificateErrorDetail string `json:"certificateErrorDetail,omitempty"`

	// Certificate related properties which may provide additional information about the given certificate.
	// Required: true
	Properties map[string]string `json:"properties"`
}


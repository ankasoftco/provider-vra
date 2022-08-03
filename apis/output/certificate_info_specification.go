package models 
type CertificateInfoSpecification struct {

	// The certificate in string format.
	// Example: PEM for X509Certificate
	// Required: true
	Certificate *string `json:"certificate"`
}


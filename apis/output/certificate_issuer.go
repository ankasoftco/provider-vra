package models 
type CertificateIssuer struct {

	// A human-friendly name used as an identifier for the holding body.
	// Example: vmware.com
	CommonName string `json:"commonName,omitempty"`

	// Name of the organisation.
	// Example: VMware Inc.
	Organization string `json:"organization,omitempty"`
}


package models 
type endpointCertificate struct {
	fingerprintsField CertificateFingerprint

	issuedByField *CertificateIssuer

	issuedToField CertificateIssuedTo

	periodOfValidityField CertificateValidity
}


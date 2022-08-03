package models 
type EndpointProperties struct {

	// accept self signed certificate
	AcceptSelfSignedCertificate bool `json:"acceptSelfSignedCertificate,omitempty"`

	// certificate
	Certificate string `json:"certificate,omitempty"`

	// dc Id
	DcID string `json:"dcId,omitempty"`

	// endpoint Id
	EndpointID string `json:"endpointId,omitempty"`

	// host name
	HostName string `json:"hostName,omitempty"`

	// private key
	PrivateKey string `json:"privateKey,omitempty"`

	// private key Id
	PrivateKeyID string `json:"privateKeyId,omitempty"`

	// server
	Server string `json:"server,omitempty"`

	// service account Id
	ServiceAccountID string `json:"serviceAccountId,omitempty"`

	// url
	URL string `json:"url,omitempty"`
}


package models 
type X509Certificate struct {

	// encoded
	// Format: byte
	Encoded strfmt.Base64 `json:"encoded,omitempty"`

	issuerDNField Principal

	// not after
	// Format: date-time
	NotAfter strfmt.DateTime `json:"notAfter,omitempty"`

	// not before
	// Format: date-time
	NotBefore strfmt.DateTime `json:"notBefore,omitempty"`

	// public key
	PublicKey *PublicKey `json:"publicKey,omitempty"`

	// serial number
	SerialNumber int64 `json:"serialNumber,omitempty"`

	// sig alg name
	SigAlgName string `json:"sigAlgName,omitempty"`

	// sig alg o ID
	SigAlgOID string `json:"sigAlgOID,omitempty"`

	// sig alg params
	// Format: byte
	SigAlgParams strfmt.Base64 `json:"sigAlgParams,omitempty"`

	subjectDNField Principal

	// version
	Version int32 `json:"version,omitempty"`
}

// UnmarshalJSON unmarshals this object with a polymorphic type from a JSON structure
func (m *X509Certificate) UnmarshalJSON(raw []byte) error {
	var data struct {
		Encoded strfmt.Base64 `json:"encoded,omitempty"`

		IssuerDN json.RawMessage `json:"issuerDN,omitempty"`

		NotAfter strfmt.DateTime `json:"notAfter,omitempty"`

		NotBefore strfmt.DateTime `json:"notBefore,omitempty"`

		PublicKey *PublicKey `json:"publicKey,omitempty"`

		SerialNumber int64 `json:"serialNumber,omitempty"`

		SigAlgName string `json:"sigAlgName,omitempty"`

		SigAlgOID string `json:"sigAlgOID,omitempty"`

		SigAlgParams strfmt.Base64 `json:"sigAlgParams,omitempty"`

		SubjectDN json.RawMessage `json:"subjectDN,omitempty"`

		Version int32 `json:"version,omitempty"`
	}

// MarshalJSON marshals this object with a polymorphic type to a JSON structure
func (m X509Certificate) MarshalJSON() ([]byte, error) {
	var b1, b2, b3 []byte
	var err error
	b1, err = json.Marshal(struct {
		Encoded strfmt.Base64 `json:"encoded,omitempty"`

		NotAfter strfmt.DateTime `json:"notAfter,omitempty"`

		NotBefore strfmt.DateTime `json:"notBefore,omitempty"`

		PublicKey *PublicKey `json:"publicKey,omitempty"`

		SerialNumber int64 `json:"serialNumber,omitempty"`

		SigAlgName string `json:"sigAlgName,omitempty"`

		SigAlgOID string `json:"sigAlgOID,omitempty"`

		SigAlgParams strfmt.Base64 `json:"sigAlgParams,omitempty"`

		Version int32 `json:"version,omitempty"`
	}{

		Encoded: m.Encoded,

		NotAfter: m.NotAfter,

		NotBefore: m.NotBefore,

		PublicKey: m.PublicKey,

		SerialNumber: m.SerialNumber,

		SigAlgName: m.SigAlgName,

		SigAlgOID: m.SigAlgOID,

		SigAlgParams: m.SigAlgParams,

		Version: m.Version,
	})
	if err != nil {
		return nil, err
	}


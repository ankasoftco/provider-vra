package models 
type TagState struct {

	// deleted
	Deleted bool `json:"deleted,omitempty"`

	// document auth principal link
	DocumentAuthPrincipalLink string `json:"documentAuthPrincipalLink,omitempty"`

	// document description
	DocumentDescription *ServiceDocumentDescription `json:"documentDescription,omitempty"`

	// document epoch
	DocumentEpoch int64 `json:"documentEpoch,omitempty"`

	// document expiration time micros
	DocumentExpirationTimeMicros int64 `json:"documentExpirationTimeMicros,omitempty"`

	// document kind
	DocumentKind string `json:"documentKind,omitempty"`

	// document owner
	DocumentOwner string `json:"documentOwner,omitempty"`

	// document self link
	DocumentSelfLink string `json:"documentSelfLink,omitempty"`

	// document source link
	DocumentSourceLink string `json:"documentSourceLink,omitempty"`

	// document transaction Id
	DocumentTransactionID string `json:"documentTransactionId,omitempty"`

	// document update action
	DocumentUpdateAction string `json:"documentUpdateAction,omitempty"`

	// document update time micros
	DocumentUpdateTimeMicros int64 `json:"documentUpdateTimeMicros,omitempty"`

	// document version
	DocumentVersion int64 `json:"documentVersion,omitempty"`

	// external
	External bool `json:"external,omitempty"`

	// key
	Key string `json:"key,omitempty"`

	// origins
	Origins []string `json:"origins"`

	// tenant links
	TenantLinks []string `json:"tenantLinks"`

	// value
	Value string `json:"value,omitempty"`
}


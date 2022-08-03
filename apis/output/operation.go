package models 
type Operation struct {

	// action
	// Enum: [GET POST PATCH PUT DELETE OPTIONS]
	Action string `json:"action,omitempty"`

	// authorization context
	AuthorizationContext *AuthorizationContext `json:"authorizationContext,omitempty"`

	// body raw
	BodyRaw interface{} `json:"bodyRaw,omitempty"`

	// commit
	Commit bool `json:"commit,omitempty"`

	// completion
	Completion CompletionHandler `json:"completion,omitempty"`

	// connection sharing
	ConnectionSharing bool `json:"connectionSharing,omitempty"`

	// connection tag
	ConnectionTag string `json:"connectionTag,omitempty"`

	// content length
	ContentLength int64 `json:"contentLength,omitempty"`

	// content type
	ContentType string `json:"contentType,omitempty"`

	// context Id
	ContextID string `json:"contextId,omitempty"`

	// cookies
	Cookies map[string]string `json:"cookies,omitempty"`

	// error response body
	ErrorResponseBody *ServiceErrorResponse `json:"errorResponseBody,omitempty"`

	// expiration micros utc
	ExpirationMicrosUtc int64 `json:"expirationMicrosUtc,omitempty"`

	// failure logging disabled
	FailureLoggingDisabled bool `json:"failureLoggingDisabled,omitempty"`

	// forwarded
	Forwarded bool `json:"forwarded,omitempty"`

	// forwarding disabled
	ForwardingDisabled bool `json:"forwardingDisabled,omitempty"`

	// from replication
	FromReplication bool `json:"fromReplication,omitempty"`

	// id
	ID int64 `json:"id,omitempty"`

	// keep alive
	KeepAlive bool `json:"keepAlive,omitempty"`

	// notification
	Notification bool `json:"notification,omitempty"`

	// notification disabled
	NotificationDisabled bool `json:"notificationDisabled,omitempty"`

	// options
	Options []string `json:"options"`

	// peer certificate chain
	PeerCertificateChain []*X509Certificate `json:"peerCertificateChain"`

	peerPrincipalField Principal

	// referer
	Referer *URI `json:"referer,omitempty"`

	// referer as string
	RefererAsString string `json:"refererAsString,omitempty"`

	// remote
	Remote bool `json:"remote,omitempty"`

	// replication disabled
	ReplicationDisabled bool `json:"replicationDisabled,omitempty"`

	// request headers
	RequestHeaders map[string]string `json:"requestHeaders,omitempty"`

	// response headers
	ResponseHeaders map[string]string `json:"responseHeaders,omitempty"`

	// retries remaining
	RetriesRemaining int32 `json:"retriesRemaining,omitempty"`

	// retry count
	RetryCount int32 `json:"retryCount,omitempty"`

	// socket context
	SocketContext *SocketContext `json:"socketContext,omitempty"`

	// status code
	StatusCode int32 `json:"statusCode,omitempty"`

	// synchronize
	Synchronize bool `json:"synchronize,omitempty"`

	// synchronize owner
	SynchronizeOwner bool `json:"synchronizeOwner,omitempty"`

	// synchronize peer
	SynchronizePeer bool `json:"synchronizePeer,omitempty"`

	// target replicated
	TargetReplicated bool `json:"targetReplicated,omitempty"`

	// transaction Id
	TransactionID string `json:"transactionId,omitempty"`

	// update
	Update bool `json:"update,omitempty"`

	// uri
	URI *URI `json:"uri,omitempty"`

	// within transaction
	WithinTransaction bool `json:"withinTransaction,omitempty"`
}

// UnmarshalJSON unmarshals this object with a polymorphic type from a JSON structure
func (m *Operation) UnmarshalJSON(raw []byte) error {
	var data struct {
		Action string `json:"action,omitempty"`

		AuthorizationContext *AuthorizationContext `json:"authorizationContext,omitempty"`

		BodyRaw interface{} `json:"bodyRaw,omitempty"`

		Commit bool `json:"commit,omitempty"`

		Completion CompletionHandler `json:"completion,omitempty"`

		ConnectionSharing bool `json:"connectionSharing,omitempty"`

		ConnectionTag string `json:"connectionTag,omitempty"`

		ContentLength int64 `json:"contentLength,omitempty"`

		ContentType string `json:"contentType,omitempty"`

		ContextID string `json:"contextId,omitempty"`

		Cookies map[string]string `json:"cookies,omitempty"`

		ErrorResponseBody *ServiceErrorResponse `json:"errorResponseBody,omitempty"`

		ExpirationMicrosUtc int64 `json:"expirationMicrosUtc,omitempty"`

		FailureLoggingDisabled bool `json:"failureLoggingDisabled,omitempty"`

		Forwarded bool `json:"forwarded,omitempty"`

		ForwardingDisabled bool `json:"forwardingDisabled,omitempty"`

		FromReplication bool `json:"fromReplication,omitempty"`

		ID int64 `json:"id,omitempty"`

		KeepAlive bool `json:"keepAlive,omitempty"`

		Notification bool `json:"notification,omitempty"`

		NotificationDisabled bool `json:"notificationDisabled,omitempty"`

		Options []string `json:"options"`

		PeerCertificateChain []*X509Certificate `json:"peerCertificateChain"`

		PeerPrincipal json.RawMessage `json:"peerPrincipal,omitempty"`

		Referer *URI `json:"referer,omitempty"`

		RefererAsString string `json:"refererAsString,omitempty"`

		Remote bool `json:"remote,omitempty"`

		ReplicationDisabled bool `json:"replicationDisabled,omitempty"`

		RequestHeaders map[string]string `json:"requestHeaders,omitempty"`

		ResponseHeaders map[string]string `json:"responseHeaders,omitempty"`

		RetriesRemaining int32 `json:"retriesRemaining,omitempty"`

		RetryCount int32 `json:"retryCount,omitempty"`

		SocketContext *SocketContext `json:"socketContext,omitempty"`

		StatusCode int32 `json:"statusCode,omitempty"`

		Synchronize bool `json:"synchronize,omitempty"`

		SynchronizeOwner bool `json:"synchronizeOwner,omitempty"`

		SynchronizePeer bool `json:"synchronizePeer,omitempty"`

		TargetReplicated bool `json:"targetReplicated,omitempty"`

		TransactionID string `json:"transactionId,omitempty"`

		Update bool `json:"update,omitempty"`

		URI *URI `json:"uri,omitempty"`

		WithinTransaction bool `json:"withinTransaction,omitempty"`
	}

// MarshalJSON marshals this object with a polymorphic type to a JSON structure
func (m Operation) MarshalJSON() ([]byte, error) {
	var b1, b2, b3 []byte
	var err error
	b1, err = json.Marshal(struct {
		Action string `json:"action,omitempty"`

		AuthorizationContext *AuthorizationContext `json:"authorizationContext,omitempty"`

		BodyRaw interface{} `json:"bodyRaw,omitempty"`

		Commit bool `json:"commit,omitempty"`

		Completion CompletionHandler `json:"completion,omitempty"`

		ConnectionSharing bool `json:"connectionSharing,omitempty"`

		ConnectionTag string `json:"connectionTag,omitempty"`

		ContentLength int64 `json:"contentLength,omitempty"`

		ContentType string `json:"contentType,omitempty"`

		ContextID string `json:"contextId,omitempty"`

		Cookies map[string]string `json:"cookies,omitempty"`

		ErrorResponseBody *ServiceErrorResponse `json:"errorResponseBody,omitempty"`

		ExpirationMicrosUtc int64 `json:"expirationMicrosUtc,omitempty"`

		FailureLoggingDisabled bool `json:"failureLoggingDisabled,omitempty"`

		Forwarded bool `json:"forwarded,omitempty"`

		ForwardingDisabled bool `json:"forwardingDisabled,omitempty"`

		FromReplication bool `json:"fromReplication,omitempty"`

		ID int64 `json:"id,omitempty"`

		KeepAlive bool `json:"keepAlive,omitempty"`

		Notification bool `json:"notification,omitempty"`

		NotificationDisabled bool `json:"notificationDisabled,omitempty"`

		Options []string `json:"options"`

		PeerCertificateChain []*X509Certificate `json:"peerCertificateChain"`

		Referer *URI `json:"referer,omitempty"`

		RefererAsString string `json:"refererAsString,omitempty"`

		Remote bool `json:"remote,omitempty"`

		ReplicationDisabled bool `json:"replicationDisabled,omitempty"`

		RequestHeaders map[string]string `json:"requestHeaders,omitempty"`

		ResponseHeaders map[string]string `json:"responseHeaders,omitempty"`

		RetriesRemaining int32 `json:"retriesRemaining,omitempty"`

		RetryCount int32 `json:"retryCount,omitempty"`

		SocketContext *SocketContext `json:"socketContext,omitempty"`

		StatusCode int32 `json:"statusCode,omitempty"`

		Synchronize bool `json:"synchronize,omitempty"`

		SynchronizeOwner bool `json:"synchronizeOwner,omitempty"`

		SynchronizePeer bool `json:"synchronizePeer,omitempty"`

		TargetReplicated bool `json:"targetReplicated,omitempty"`

		TransactionID string `json:"transactionId,omitempty"`

		Update bool `json:"update,omitempty"`

		URI *URI `json:"uri,omitempty"`

		WithinTransaction bool `json:"withinTransaction,omitempty"`
	}{

		Action: m.Action,

		AuthorizationContext: m.AuthorizationContext,

		BodyRaw: m.BodyRaw,

		Commit: m.Commit,

		Completion: m.Completion,

		ConnectionSharing: m.ConnectionSharing,

		ConnectionTag: m.ConnectionTag,

		ContentLength: m.ContentLength,

		ContentType: m.ContentType,

		ContextID: m.ContextID,

		Cookies: m.Cookies,

		ErrorResponseBody: m.ErrorResponseBody,

		ExpirationMicrosUtc: m.ExpirationMicrosUtc,

		FailureLoggingDisabled: m.FailureLoggingDisabled,

		Forwarded: m.Forwarded,

		ForwardingDisabled: m.ForwardingDisabled,

		FromReplication: m.FromReplication,

		ID: m.ID,

		KeepAlive: m.KeepAlive,

		Notification: m.Notification,

		NotificationDisabled: m.NotificationDisabled,

		Options: m.Options,

		PeerCertificateChain: m.PeerCertificateChain,

		Referer: m.Referer,

		RefererAsString: m.RefererAsString,

		Remote: m.Remote,

		ReplicationDisabled: m.ReplicationDisabled,

		RequestHeaders: m.RequestHeaders,

		ResponseHeaders: m.ResponseHeaders,

		RetriesRemaining: m.RetriesRemaining,

		RetryCount: m.RetryCount,

		SocketContext: m.SocketContext,

		StatusCode: m.StatusCode,

		Synchronize: m.Synchronize,

		SynchronizeOwner: m.SynchronizeOwner,

		SynchronizePeer: m.SynchronizePeer,

		TargetReplicated: m.TargetReplicated,

		TransactionID: m.TransactionID,

		Update: m.Update,

		URI: m.URI,

		WithinTransaction: m.WithinTransaction,
	})
	if err != nil {
		return nil, err
	}


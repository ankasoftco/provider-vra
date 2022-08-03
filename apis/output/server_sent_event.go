package models 
type ServerSentEvent struct {

	// completion callback
	CompletionCallback *Throwable `json:"completionCallback,omitempty"`

	// data
	Data string `json:"data,omitempty"`

	// event
	Event string `json:"event,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// retry
	Retry int64 `json:"retry,omitempty"`
}


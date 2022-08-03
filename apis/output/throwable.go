package models 
type Throwable struct {

	// cause
	Cause *Throwable `json:"cause,omitempty"`

	// stack trace
	StackTrace []*StackTraceElement `json:"stackTrace"`
}


package models 
type User struct {

	// The email of the user or name of the group.
	// Example: administrator@vmware.com
	// Required: true
	Email *string `json:"email"`

	// Type of the principal. Currently supported 'user' (default) and 'group'.
	// Example: user
	Type string `json:"type,omitempty"`
}


package models 
type AuthorizationContext struct {

	// claims
	Claims *Claims `json:"claims,omitempty"`

	// guest user
	GuestUser bool `json:"guestUser,omitempty"`

	// system user
	SystemUser bool `json:"systemUser,omitempty"`

	// token
	Token string `json:"token,omitempty"`
}


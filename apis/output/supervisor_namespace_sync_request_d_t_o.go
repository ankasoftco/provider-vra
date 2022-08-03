package models 
type SupervisorNamespaceSyncRequestDTO struct {

	// edit groups
	EditGroups string `json:"editGroups,omitempty"`

	// edit users
	EditUsers string `json:"editUsers,omitempty"`

	// view groups
	ViewGroups string `json:"viewGroups,omitempty"`

	// view users
	ViewUsers string `json:"viewUsers,omitempty"`
}


package models 
type DeploymentReference struct {

	// description
	Description string `json:"description,omitempty"`

	// icon Id
	IconID string `json:"iconId,omitempty"`

	// id
	// Format: uuid
	ID strfmt.UUID `json:"id,omitempty"`

	// lease expire at
	// Format: date-time
	LeaseExpireAt strfmt.DateTime `json:"leaseExpireAt,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// owned by
	OwnedBy string `json:"ownedBy,omitempty"`
}


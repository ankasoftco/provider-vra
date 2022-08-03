package models 
type MeteringPolicyAssignment struct {

	// Creation time
	// Format: date-time
	CreatedAt strfmt.DateTime `json:"createdAt,omitempty"`

	// Pricing card assigned entity id
	EntityID string `json:"entityId,omitempty"`

	// Pricing card assigned entity name
	EntityName string `json:"entityName,omitempty"`

	// Pricing card assigned entity type
	// Enum: [ALL PROJECT CLOUDZONE]
	EntityType string `json:"entityType,omitempty"`

	// Id of the pricingCardAssignment
	// Format: uuid
	ID strfmt.UUID `json:"id,omitempty"`

	// Updated time
	// Format: date-time
	LastUpdatedAt strfmt.DateTime `json:"lastUpdatedAt,omitempty"`

	// OrgId of the pricingCardAssignment
	OrgID string `json:"orgId,omitempty"`

	// Pricing card id
	// Format: uuid
	PricingCardID strfmt.UUID `json:"pricingCardId,omitempty"`

	// Pricing card name
	PricingCardName string `json:"pricingCardName,omitempty"`
}


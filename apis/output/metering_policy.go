package models 
type MeteringPolicy struct {

	// charge model
	// Enum: [PAY_AS_YOU_GO]
	ChargeModel string `json:"chargeModel,omitempty"`

	// created at
	// Format: date-time
	CreatedAt strfmt.DateTime `json:"createdAt,omitempty"`

	// created by
	CreatedBy string `json:"createdBy,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// fixed price
	FixedPrice *FixedPrice `json:"fixedPrice,omitempty"`

	// id
	// Format: uuid
	ID strfmt.UUID `json:"id,omitempty"`

	// last updated at
	// Format: date-time
	LastUpdatedAt strfmt.DateTime `json:"lastUpdatedAt,omitempty"`

	// metering items
	MeteringItems []*MeteringItem `json:"meteringItems"`

	// name
	Name string `json:"name,omitempty"`

	// named metering items
	NamedMeteringItems []*NamedMeteringItem `json:"namedMeteringItems"`

	// one time metering items
	OneTimeMeteringItems []*OneTimeMeteringItem `json:"oneTimeMeteringItems"`

	// org Id
	OrgID string `json:"orgId,omitempty"`

	// Assignment count, assignment entity type
	PricingCardAssignmentInfo *MeteringPolicyAssignmentInfo `json:"pricingCardAssignmentInfo,omitempty"`

	// tag based metering items
	TagBasedMeteringItems []*TagBasedMeteringItem `json:"tagBasedMeteringItems"`

	// tag based one time metering items
	TagBasedOneTimeMeteringItems []*TagBasedOneTimeMeteringItem `json:"tagBasedOneTimeMeteringItems"`

	// tag based rate factor items
	TagBasedRateFactorItems []*TagBasedRateFactorItem `json:"tagBasedRateFactorItems"`
}


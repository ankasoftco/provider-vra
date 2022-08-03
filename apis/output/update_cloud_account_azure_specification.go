package models 
type UpdateCloudAccountAzureSpecification struct {

	// Azure Client Application ID
	// Example: 3287dd6e-76d8-41b7-9856-2584969e7739
	// Required: true
	ClientApplicationID *string `json:"clientApplicationId"`

	// Azure Client Application Secret Key
	// Example: GDfdasDasdASFas321das32cas2x3dsXCSA76xdcasg=
	// Required: true
	ClientApplicationSecretKey *string `json:"clientApplicationSecretKey"`

	// Create default cloud zones for the enabled regions.
	// Example: true
	CreateDefaultZones bool `json:"createDefaultZones,omitempty"`

	// A human-friendly description.
	Description string `json:"description,omitempty"`

	// A human-friendly name used as an identifier in APIs that support this option.
	Name string `json:"name,omitempty"`

	// A set of regions to enable provisioning on.Refer to /iaas/api/cloud-accounts/region-enumeration.
	// Example: [{ \"name\": \"East Asia\",\"externalRegionId\": \"eastasia\"}]
	Regions []*RegionSpecification `json:"regions"`

	// Azure Subscribtion ID
	// Example: 064865b2-e914-4717-b415-8806d17948f7
	// Required: true
	SubscriptionID *string `json:"subscriptionId"`

	// A set of tag keys and optional values to set on the Cloud Account
	// Example: [ { \"key\" : \"env\", \"value\": \"dev\" } ]
	Tags []*Tag `json:"tags"`

	// Azure Tenant ID
	// Example: 9a13d920-4691-4e2d-b5d5-9c4c1279bc9a
	// Required: true
	TenantID *string `json:"tenantId"`
}


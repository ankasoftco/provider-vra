package models 
type CloudAccountAwsSpecification struct {

	// Aws Access key ID
	// Example: ACDC55DB4MFH6ADG75KK
	// Required: true
	AccessKeyID *string `json:"accessKeyId"`

	// Create default cloud zones for the enabled regions.
	// Example: true
	CreateDefaultZones bool `json:"createDefaultZones,omitempty"`

	// A human-friendly description.
	Description string `json:"description,omitempty"`

	// A human-friendly name used as an identifier in APIs that support this option.
	// Required: true
	Name *string `json:"name"`

	// A set of regions to enable provisioning on.Refer to /iaas/api/cloud-accounts/region-enumeration.
	// Example: [{ \"name\": \"eu-west-1\",\"externalRegionId\": \"eu-west-1\"}]
	// Required: true
	Regions []*RegionSpecification `json:"regions"`

	// Aws Secret Access Key
	// Example: gfsScK345sGGaVdds222dasdfDDSSasdfdsa34fS
	// Required: true
	SecretAccessKey *string `json:"secretAccessKey"`

	// A set of tag keys and optional values to set on the Cloud Account
	// Example: [ { \"key\" : \"env\", \"value\": \"dev\" } ]
	Tags []*Tag `json:"tags"`
}


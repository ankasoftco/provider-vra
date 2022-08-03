package models 
type ConfigurationProperty struct {

	// The key of the property.
	// Required: true
	// Enum: [SESSION_TIMEOUT_DURATION_MINUTES RELEASE_IPADDRESS_PERIOD_MINUTES NSXT_RETRY_DURATION_MINUTES]
	Key *string `json:"key"`

	// The value of the property.
	// Required: true
	Value *string `json:"value"`
}


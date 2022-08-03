package models 
type DataCollectorRegistration struct {

	// A registration key for the data collector
	// Example: eyJyZWdpc3RyYXRpb25VcmwiOiJodHRwczovL2Fw
	// Required: true
	Key *string `json:"key"`

	// Data collector OVA Link
	// Example: https://ci-data-collector.s3.amazonaws.com/VMware-Cloud-Services-Data-Collector.ova
	// Required: true
	OvaLink *string `json:"ovaLink"`
}


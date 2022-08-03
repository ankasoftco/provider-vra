package models 
type DataCollector struct {

	// Data collector identifier
	// Example: d5316b00-f3b8-4895-9e9a-c4b98649c2ca
	// Required: true
	DcID *string `json:"dcId"`

	// Data collector host name
	// Example: dc1-lnd.mycompany.com
	// Required: true
	HostName *string `json:"hostName"`

	// Ip Address of the data collector VM
	// Example: 10.0.0.1
	// Required: true
	IPAddress *string `json:"ipAddress"`

	// Data collector name
	// Example: Datacollector1
	// Required: true
	Name *string `json:"name"`

	// Current status of the data collector
	// Example: ACTIVE, INACTIVE
	// Required: true
	Status *string `json:"status"`
}


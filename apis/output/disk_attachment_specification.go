package models 
type DiskAttachmentSpecification struct {

	// The id of the existing block device
	// Example: 1298765
	// Required: true
	BlockDeviceID *string `json:"blockDeviceId"`

	// A human-friendly description.
	Description string `json:"description,omitempty"`

	// Disk Attachment specific properties
	// Example: { \"scsiController\": \"SCSI_Controller_0\",\"unitNumber\" : \"2\" }
	DiskAttachmentProperties map[string]string `json:"diskAttachmentProperties,omitempty"`

	// A human-friendly name used as an identifier in APIs that support this option.
	Name string `json:"name,omitempty"`

	// Deprecated: The SCSI controller to be assigned
	// Example: SCSI_Controller_0, SCSI_Controller_1, SCSI_Controller_2, SCSI_Controller_3
	ScsiController string `json:"scsiController,omitempty"`

	// Deprecated: The Unit Number to be assigned
	// Example: 0
	UnitNumber string `json:"unitNumber,omitempty"`
}


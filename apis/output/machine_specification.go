package models 
type MachineSpecification struct {

	// A valid cloud config data in json-escaped yaml syntax
	BootConfig *MachineBootConfig `json:"bootConfig,omitempty"`

	// A set of settings that specify how the provided Boot Config should be handled
	BootConfigSettings *MachineBootConfigSettings `json:"bootConfigSettings,omitempty"`

	// Constraints that are used to drive placement policies for the virtual machine that is produced from this specification. Constraint expressions are matched against tags on existing placement targets.
	// Example: [{\"mandatory\" : \"true\", \"expression\": \"environment\":\"prod\"}, {\"mandatory\" : \"false\", \"expression\": \"pci\"}]
	Constraints []*Constraint `json:"constraints"`

	// Additional custom properties that may be used to extend this resource.
	CustomProperties map[string]string `json:"customProperties,omitempty"`

	// The id of the deployment that is associated with this resource
	// Example: 123e4567-e89b-12d3-a456-426655440000
	DeploymentID string `json:"deploymentId,omitempty"`

	// Describes machine within the scope of your organization and is not propagated to the cloud
	Description string `json:"description,omitempty"`

	// A set of disk specifications for this machine.
	Disks []*DiskAttachmentSpecification `json:"disks"`

	// Flavor of machine instance.
	// Example: small, medium, large
	// Required: true
	Flavor *string `json:"flavor"`

	// Provider specific flavor reference. Valid if no flavor property is provided
	// Example: t2.micro
	// Required: true
	FlavorRef *string `json:"flavorRef"`

	// Type of image used for this machine.
	// Example: vmware-gold-master, ubuntu-latest, rhel-compliant, windows
	// Required: true
	Image *string `json:"image"`

	// Constraints that are used to drive placement policies for the image disk. Constraint expressions are matched against tags on existing placement targets.
	// Example: [{\"mandatory\" : \"true\", \"expression\": \"environment:prod\"}, {\"mandatory\" : \"false\", \"expression\": \"pci\"}]
	ImageDiskConstraints []*Constraint `json:"imageDiskConstraints"`

	// Direct image reference used for this machine (name, path, location, uri, etc.). Valid if no image property is provided
	// Example: ami-f6795a8c
	// Required: true
	ImageRef *string `json:"imageRef"`

	// Number of machines to provision - default 1.
	// Example: 3
	MachineCount int32 `json:"machineCount,omitempty"`

	// A human-friendly name used as an identifier in APIs that support this option.
	// Required: true
	Name *string `json:"name"`

	// A set of network interface controller specifications for this machine. If not specified, then a default network connection will be created.
	Nics []*NetworkInterfaceSpecification `json:"nics"`

	// The id of the project the current user belongs to.
	// Example: e058
	// Required: true
	ProjectID *string `json:"projectId"`

	// Settings to remotely connect to the provisioned machine, by public/private key pair or username/password authentication. AWS and vSphere support key pair. Azure supports key pair or username/password.
	RemoteAccess *RemoteAccessSpecification `json:"remoteAccess,omitempty"`

	// Settings to apply salt configuration on the provisioned machine.
	SaltConfiguration *SaltConfiguration `json:"saltConfiguration,omitempty"`

	// A set of tag keys and optional values that should be set on any resource that is produced from this specification.
	// Example: [ { \"key\" : \"ownedBy\", \"value\": \"Rainpole\" } ]
	Tags []*Tag `json:"tags"`
}


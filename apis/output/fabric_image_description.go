package models 
type FabricImageDescription struct {

	// Cloud config for this image. This cloud config will be merged during provisioning with other cloud configurations such as the bootConfig provided in MachineSpecification.
	// Example: runcmd:\n  - [\"mkdir\", \"/imageFolder\"]
	CloudConfig string `json:"cloudConfig,omitempty"`

	// Constraints that are used to drive placement policies for the image that is produced from this mapping.Constraint expressions are matched against tags on existing placement targets.
	// Example: [{\"mandatory\" : \"true\", \"expression\": \"environment\":\"prod\"}, {\"mandatory\" : \"false\", \"expression\": \"pci\"}]
	Constraints []*Constraint `json:"constraints"`

	// The id of the fabric image
	// Example: 9e49
	ID string `json:"id,omitempty"`

	// Fabric image name. Valid if id not provided.
	// Example: ami-ubuntu-16.04-1.9.1-00-1516139717
	Name string `json:"name,omitempty"`
}


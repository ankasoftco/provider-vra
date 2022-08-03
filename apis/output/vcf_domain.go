package models 
type VcfDomain struct {

	// id
	ID string `json:"id,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// nsx resource
	NsxResource *NsxAccount `json:"nsxResource,omitempty"`

	// status
	// Enum: [ACTIVE CONFIGURED NOT_CONFIGURED NOT_ACTIVE]
	Status string `json:"status,omitempty"`

	// type
	Type string `json:"type,omitempty"`

	// vsphere
	Vsphere *VsphereAccount `json:"vsphere,omitempty"`
}


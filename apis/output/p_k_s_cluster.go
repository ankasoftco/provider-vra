package models 
type PKSCluster struct {

	// hostname address
	HostnameAddress string `json:"hostnameAddress,omitempty"`

	// ip address
	IPAddress string `json:"ipAddress,omitempty"`

	// kubernetes master ips
	KubernetesMasterIps []string `json:"kubernetes_master_ips"`

	// last action
	LastAction string `json:"last_action,omitempty"`

	// last action description
	LastActionDescription string `json:"last_action_description,omitempty"`

	// last action state
	LastActionState string `json:"last_action_state,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// parameters
	Parameters interface{} `json:"parameters,omitempty"`

	// plan name
	PlanName string `json:"plan_name,omitempty"`

	// uuid
	UUID string `json:"uuid,omitempty"`
}


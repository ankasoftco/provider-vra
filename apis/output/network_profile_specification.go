package models 
type NetworkProfileSpecification struct {

	// Additional properties that may be used to extend the Network Profile object that is produced from this specification.  For isolationType security group, datastoreId identifies the Compute Resource Edge datastore. computeCluster and resourcePoolId identify the Compute Resource Edge cluster. For isolationType subnet, distributedLogicalRouterStateLink identifies the on-demand network distributed local router.  onDemandNetworkIPAssignmentType identifies the on-demand network IP range assignment type static, dynamic, or mixed.
	// Example: { \"resourcePoolId\" : \"resource-pool-1\", \"datastoreId\" : \"StoragePod:group-p87839\", \"computeCluster\" : \"/resources/compute/1234\", \"distributedLogicalRouterStateLink\" : \"/resources/routers/1234\", \"onDemandNetworkIPAssignmentType\" : \"dynamic\"}
	CustomProperties map[string]string `json:"customProperties,omitempty"`

	// A human-friendly description.
	Description string `json:"description,omitempty"`

	// List of external IP blocks coming from an external IPAM provider that can be used to create subnetworks inside them
	// Example: [\"3e2bb9bc-6a6a-11ea-bc55-0242ac130003\"]
	ExternalIPBlockIds []string `json:"externalIpBlockIds"`

	// A list of fabric network Ids which are assigned to the network profile.
	// Example: [ \"6543\" ]
	FabricNetworkIds []string `json:"fabricNetworkIds"`

	// The CIDR prefix length to be used for the isolated networks that are created with the network profile.
	// Example: 24
	IsolatedNetworkCIDRPrefix int32 `json:"isolatedNetworkCIDRPrefix,omitempty"`

	// The Id of the fabric network used for outbound access.
	// Example: 1234
	IsolationExternalFabricNetworkID string `json:"isolationExternalFabricNetworkId,omitempty"`

	// CIDR of the isolation network domain.
	// Example: 10.10.10.0/24
	IsolationNetworkDomainCIDR string `json:"isolationNetworkDomainCIDR,omitempty"`

	// The Id of the network domain used for creating isolated networks.
	// Example: 1234
	IsolationNetworkDomainID string `json:"isolationNetworkDomainId,omitempty"`

	// Specifies the isolation type e.g. none, subnet or security group
	// Example: SUBNET
	// Enum: [NONE SUBNET SECURITY_GROUP]
	IsolationType string `json:"isolationType,omitempty"`

	// A list of load balancers which are assigned to the network profile.
	// Example: [ \"6545\" ]
	LoadBalancerIds []string `json:"loadBalancerIds"`

	// A human-friendly name used as an identifier in APIs that support this option.
	// Required: true
	Name *string `json:"name"`

	// The Id of the region for which this profile is created
	// Example: 9e49
	// Required: true
	RegionID *string `json:"regionId"`

	// A list of security group Ids which are assigned to the network profile.
	// Example: [ \"6545\" ]
	SecurityGroupIds []string `json:"securityGroupIds"`

	// A set of tag keys and optional values that should be set on any resource that is produced from this specification.
	// Example: [ { \"key\" : \"dev\", \"value\": \"hard\" } ]
	Tags []*Tag `json:"tags"`
}


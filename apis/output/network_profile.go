package models 
type NetworkProfile struct {

	// HATEOAS of the entity
	// Required: true
	Links map[string]Href `json:"_links"`

	// Id of the cloud account this profile belongs to.
	// Example: [9e49]
	CloudAccountID string `json:"cloudAccountId,omitempty"`

	// Date when the entity was created. The date is in ISO 8601 and UTC.
	// Example: 2012-09-27
	CreatedAt string `json:"createdAt,omitempty"`

	// Additional properties that may be used to extend the Network Profile object that is produced from this specification.  For isolationType security group, datastoreId identifies the Compute Resource Edge datastore. computeCluster and resourcePoolId identify the Compute Resource Edge cluster. For isolationType subnet, distributedLogicalRouterStateLink identifies the on-demand network distributed local router.  onDemandNetworkIPAssignmentType identifies the on-demand network IP range assignment type static, dynamic, or mixed.
	// Example: { \"resourcePoolId\" : \"resource-pool-1\", \"datastoreId\" : \"StoragePod:group-p87839\", \"computeCluster\" : \"/resources/compute/1234\", \"distributedLogicalRouterStateLink\" : \"/resources/routers/1234\", \"onDemandNetworkIPAssignmentType\" : \"dynamic\"}
	CustomProperties map[string]string `json:"customProperties,omitempty"`

	// A human-friendly description.
	// Example: my-description
	Description string `json:"description,omitempty"`

	// The id of the region for which this profile is defined
	// Example: us-east-1
	ExternalRegionID string `json:"externalRegionId,omitempty"`

	// The id of this resource instance
	// Example: 9e49
	// Required: true
	ID *string `json:"id"`

	// The CIDR prefix length to be used for the isolated networks that are created with the network profile.
	IsolatedNetworkCIDRPrefix int32 `json:"isolatedNetworkCIDRPrefix,omitempty"`

	// CIDR of the isolation network domain.
	IsolationNetworkDomainCIDR string `json:"isolationNetworkDomainCIDR,omitempty"`

	// Specifies the isolation type e.g. none, subnet or security group
	// Enum: [NONE SUBNET SECURITY_GROUP]
	IsolationType string `json:"isolationType,omitempty"`

	// A human-friendly name used as an identifier in APIs that support this option.
	// Example: my-name
	Name string `json:"name,omitempty"`

	// The id of the organization this entity belongs to.
	// Example: 42413b31-1716-477e-9a88-9dc1c3cb1cdf
	OrgID string `json:"orgId,omitempty"`

	// Email of the user that owns the entity.
	// Example: csp@vmware.com
	Owner string `json:"owner,omitempty"`

	// A set of tag keys and optional values that were set on this Network Profile.
	// Example: [ { \"key\" : \"ownedBy\", \"value\": \"Rainpole\" } ]
	Tags []*Tag `json:"tags"`

	// Date when the entity was last updated. The date is ISO 8601 and UTC.
	// Example: 2012-09-27
	UpdatedAt string `json:"updatedAt,omitempty"`
}


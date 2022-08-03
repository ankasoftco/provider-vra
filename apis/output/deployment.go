package models 
type Deployment struct {

	// Expanded deployment blueprint
	Blueprint *ResourceReference `json:"blueprint,omitempty"`

	// Deployment blueprint id
	BlueprintID string `json:"blueprintId,omitempty"`

	// Deployment blueprint version
	BlueprintVersion string `json:"blueprintVersion,omitempty"`

	// Expanded deployment catalog
	Catalog *ResourceReference `json:"catalog,omitempty"`

	// Deployment catalog item id
	CatalogItemID string `json:"catalogItemId,omitempty"`

	// Deployment catalog version
	CatalogItemVersion string `json:"catalogItemVersion,omitempty"`

	// Creation time
	// Format: date-time
	CreatedAt strfmt.DateTime `json:"createdAt,omitempty"`

	// Created by
	CreatedBy string `json:"createdBy,omitempty"`

	// Indicates whether the deployment is deleted or not.
	Deleted bool `json:"deleted,omitempty"`

	// Description of the deployment
	Description string `json:"description,omitempty"`

	// Expense associated with the deployment.
	Expense *Expense `json:"expense,omitempty"`

	// Deployment icon id
	IconID string `json:"iconId,omitempty"`

	// Id of the deployment
	// Format: uuid
	ID strfmt.UUID `json:"id,omitempty"`

	// The inputs that were used to request this deployment
	Inputs interface{} `json:"inputs,omitempty"`

	// Last request
	LastRequest *Request `json:"lastRequest,omitempty"`

	// Update time
	// Format: date-time
	LastUpdatedAt strfmt.DateTime `json:"lastUpdatedAt,omitempty"`

	// Updated by
	LastUpdatedBy string `json:"lastUpdatedBy,omitempty"`

	// Lease expiration time
	// Format: date-time
	LeaseExpireAt strfmt.DateTime `json:"leaseExpireAt,omitempty"`

	// Name of the deployment
	// Required: true
	Name *string `json:"name"`

	// org Id
	OrgID string `json:"orgId,omitempty"`

	// Owned by
	OwnedBy string `json:"ownedBy,omitempty"`

	// Expanded deployment project
	Project *ResourceReference `json:"project,omitempty"`

	// Deployment project id
	ProjectID string `json:"projectId,omitempty"`

	// Expanded resources for the deployment. Content of this property will not be maintained backward compatible
	Resources []*DeploymentResource `json:"resources"`

	// Represents the status of deployment with respect to its life cycle operations - create/update/delete.
	// Enum: [CREATE_SUCCESSFUL CREATE_INPROGRESS CREATE_FAILED UPDATE_SUCCESSFUL UPDATE_INPROGRESS UPDATE_FAILED DELETE_SUCCESSFUL DELETE_INPROGRESS DELETE_FAILED]
	Status string `json:"status,omitempty"`
}


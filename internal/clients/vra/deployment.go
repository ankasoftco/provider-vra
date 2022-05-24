/*
Copyright 2021 The Crossplane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package vra

import (
	"context"
	"github.com/go-openapi/strfmt"
)

// Deployment Deployment
//
// A group of resources such as machines, network, software, etc... typically provisioned together to deliver a complete/workable application.
//
// swagger:model Deployment
type Deployment struct {
	// Deployment blueprint id
	BlueprintID string `json:"blueprintId,omitempty"`

	// Deployment blueprint version
	BlueprintVersion string `json:"blueprintVersion,omitempty"`

	// Deployment catalog item id
	CatalogItemID string `json:"catalogItemId,omitempty"`

	// Deployment catalog version
	CatalogItemVersion string `json:"catalogItemVersion,omitempty"`

	// Created by
	CreatedBy string `json:"createdBy,omitempty"`

	// Indicates whether the vra is deleted or not.
	Deleted bool `json:"deleted,omitempty"`

	// Description of the vra
	Description string `json:"description,omitempty"`

	// Deployment icon id
	IconID string `json:"iconId,omitempty"`

	// ID of the vra
	// Format: uuid
	ID strfmt.UUID `json:"id,omitempty"`

	// The inputs that were used to request this vra
	Inputs interface{} `json:"inputs,omitempty"`

	// Update time
	// Format: date-time
	LastUpdatedAt strfmt.DateTime `json:"lastUpdatedAt,omitempty"`

	// Updated by
	LastUpdatedBy string `json:"lastUpdatedBy,omitempty"`

	// Lease expiration time
	// Format: date-time
	LeaseExpireAt strfmt.DateTime `json:"leaseExpireAt,omitempty"`

	// Name of the vra
	// Required: true
	Name *string `json:"name"`

	// org Id
	OrgID string `json:"orgId,omitempty"`

	// Owned by
	OwnedBy string `json:"ownedBy,omitempty"`

	// Deployment project id
	ProjectID string `json:"projectId,omitempty"`

	// Represents the status of vra with respect to its life cycle operations - create/update/delete.
	// Enum: [CREATE_SUCCESSFUL CREATE_INPROGRESS CREATE_FAILED UPDATE_SUCCESSFUL UPDATE_INPROGRESS UPDATE_FAILED DELETE_SUCCESSFUL DELETE_INPROGRESS DELETE_FAILED]
	Status string `json:"status,omitempty"`
}

// DeploymentClient is the API for creating/listing/deleting/getting deployments
type DeploymentClient interface {
	Create(ctx context.Context, deployment Deployment) (result Deployment, err error)
	// Delete(ctx context.Context, id int) (err error)
	Get(ctx context.Context, id int) (result Deployment, err error)
	// Update(ctx context.Context, id int, vra Deployment) (result Deployment, err error)
}

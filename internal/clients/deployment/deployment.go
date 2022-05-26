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

package deployment

import (
	"context"

	"github.com/crossplane/provider-vra/apis/deployment/v1alpha1"
)

// Client is the API for creating/listing/deleting/getting deployments
type Client interface {
	Create(ctx context.Context, deployment CreateDeploymentOptions) (result ResponseCreateDeploymentOptions, err error)
	Get(ctx context.Context, id string) (result GetDeploymentOptions, err error)
}

// CreateDeploymentOptions request payload
type CreateDeploymentOptions struct {
	DeploymentName     string            `json:"deploymentName,omitempty"`
	CatalogItemID      string            `json:"catalogItemId,omitempty"`
	CatalogItemVersion string            `json:"catalogItemVersion,omitempty"`
	ProjectID          string            `json:"projectId,omitempty"`
	Reason             string            `json:"reason,omitempty"`
	Inputs             map[string]string `json:"inputs,omitempty"`
}

// GetDeploymentOptions request payload
type GetDeploymentOptions struct {
	CreatedAt   string  `json:"createdAt,omitempty"`
	Description string  `json:"description,omitempty"`
	ID          *string `json:"id"`
	Name        string  `json:"name,omitempty"`
	OrgID       string  `json:"orgId,omitempty"`
	Owner       string  `json:"owner,omitempty"`
	ProjectID   string  `json:"projectId,omitempty"`
	UpdatedAt   string  `json:"updatedAt,omitempty"`
}

// ResponseCreateDeploymentOptions response of the create a deployment with catalog item
type ResponseCreateDeploymentOptions struct {
	DeploymentID   string `json:"deploymentId,omitempty"`
	DeploymentName string `json:"deploymentName,omitempty"`
}

// GenerateCreateDeploymentOptions returns the payload for REST API Client
func GenerateCreateDeploymentOptions(p *v1alpha1.DeploymentParameters) CreateDeploymentOptions {
	deployment := CreateDeploymentOptions{
		DeploymentName: p.DeploymentName,
		CatalogItemID:  p.CatalogItemID,
		ProjectID:      p.ProjectID,
		Inputs:         p.Inputs,
	}

	if p.CatalogItemVersion != "" {
		deployment.CatalogItemVersion = p.CatalogItemVersion
	}
	if p.Reason != "" {
		deployment.Reason = p.Reason
	}

	return deployment
}

// GenerateDeploymentObservation is get api
func GenerateDeploymentObservation(dep *GetDeploymentOptions) v1alpha1.DeploymentObservation { // nolint:gocyclo
	if dep == nil {
		return v1alpha1.DeploymentObservation{}
	}

	o := v1alpha1.DeploymentObservation{
		ID:        dep.ID,
		CreatedAt: dep.CreatedAt,
	}

	return o
}

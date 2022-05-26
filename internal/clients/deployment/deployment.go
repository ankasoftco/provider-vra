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

// DeploymentClient is the API for creating/listing/deleting/getting deployments
type DeploymentClient interface {
	Create(ctx context.Context, deployment CreateDeploymentMemberOptions) (result CreateDeploymentMemberOptions, err error)
	Get(ctx context.Context, id int) (result CreateDeploymentMemberOptions, err error)
}

type CreateDeploymentMemberOptions struct {
	DeploymentName     string            `json:"deploymentName,omitempty"`
	CatalogItemID      string            `json:"catalogItemId,omitempty"`
	CatalogItemVersion string            `json:"catalogItemVersion,omitempty"`
	ProjectID          string            `json:"projectId,omitempty"`
	Reason             string            `json:"reason,omitempty"`
	Inputs             map[string]string `json:"inputs,omitempty"`
}

// GenerateCreateDeploymentOptions returns the payload for REST API Client
func GenerateCreateDeploymentOptions(p *v1alpha1.DeploymentParameters) CreateDeploymentMemberOptions {
	deployment := CreateDeploymentMemberOptions{
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

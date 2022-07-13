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
	"github.com/anka-software/go-vra"
	"github.com/anka-software/go-vra/models"

	"github.com/crossplane/provider-vra/apis/deployment/v1alpha1"
	"github.com/crossplane/provider-vra/internal/clients"
)

// Client defines vRA Deployment service operations
type Client interface {
	GetDeployment(pid interface{}, opt *vra.GetDeploymentsOptions, options ...vra.RequestOptionFunc) (*models.Deployment, *vra.Response, error)
	CreateDeployment(opt *vra.CreateDeploymentOptions, options ...vra.RequestOptionFunc) (*models.CatalogItemRequestResponse, *vra.Response, error)
	DeleteDeployment(pid interface{}, options ...vra.RequestOptionFunc) (*vra.Response, error)
}

// NewDeploymentClient returns a new vRA Deployment service
func NewDeploymentClient(cfg clients.Config) Client {
	vra := clients.NewClient(cfg)
	return vra.Deployment
}

// GenerateObservation is used to produce v1alpha1.ProjectObservation from
// vra.Deployment.
func GenerateObservation(deployment *models.Deployment) v1alpha1.DeploymentObservation { // nolint:gocyclo
	if deployment == nil {
		return v1alpha1.DeploymentObservation{}
	}

	o := v1alpha1.DeploymentObservation{
		DeploymentID:   deployment.ID,
		DeploymentName: deployment.Name,
	}

	return o
}

// GenerateCreateDeploymentOptions generates deployment creation options
func GenerateCreateDeploymentOptions(p *v1alpha1.DeploymentParameters) *vra.CreateDeploymentOptions {
	deployment := &vra.CreateDeploymentOptions{
		CatalogItemID: p.CatalogItemID,
		RequestBody: &models.CatalogItemRequest{
			DeploymentName: p.DeploymentName,
			ProjectID:      p.ProjectID,
			Inputs:         p.Inputs,
		},
	}

	if p.Reason != "" {
		deployment.RequestBody.Reason = p.Reason
	}

	if p.CatalogItemVersion != "" {
		deployment.RequestBody.Version = p.CatalogItemVersion
	}

	return deployment
}

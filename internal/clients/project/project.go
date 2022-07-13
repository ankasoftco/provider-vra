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

package project

import (
	"github.com/crossplane/provider-vra/apis/project/v1alpha1"
	"github.com/crossplane/provider-vra/internal/clients"
	"github.com/vmware/vra-sdk-go/pkg/models"

	"github.com/vmware/vra-sdk-go/pkg/client/project"
)

// NewProjectClient returns a new vRA Deployment service
func NewProjectClient(cfg clients.Config) project.ClientService {
	vra := clients.NewClient2(cfg)

	return vra.Project
}

// GenerateCreateProjectOptions generates deployment creation options
func GenerateCreateProjectOptions(p *v1alpha1.ProjectParameters) *project.CreateProjectParams {
	var params = project.NewCreateProjectParams().WithBody(
		&models.IaaSProjectSpecification{
			Name: p.Name,
		},
	)

	return params
}

func GenerateGetProjectOptions(ProjectID string) *project.GetProjectParams {
	var params = project.NewGetProjectParams().WithID(
		ProjectID,
	)

	return params
}

func GenerateDeleteProjectOptions(ProjectID string) *project.DeleteProjectParams {
	var params = project.NewDeleteProjectParams().WithID(
		ProjectID,
	)

	return params
}

// GenerateObservation is used to produce v1alpha1.ProjectObservation from
// vra.Project.
func GenerateObservation(project *project.GetProjectOK) v1alpha1.ProjectObservation { // nolint:gocyclo
	if project.Payload == nil {
		return v1alpha1.ProjectObservation{}
	}

	o := v1alpha1.ProjectObservation{
		Name:      project.Payload.Name,
		ProjectID: project.Payload.ID,
	}

	return o
}

/*
Copyright 2022 The ANKA SOFTWARE Authors.
*/

package project

import (
	"fmt"
	"unsafe"

	"github.com/crossplane/provider-vra/apis/project/v1alpha1"
	"github.com/crossplane/provider-vra/internal/clients"
	"github.com/crossplane/provider-vra/internal/clients/login"

	"github.com/vmware/vra-sdk-go/pkg/client/project"
	"github.com/vmware/vra-sdk-go/pkg/models"
)

// NewProjectClient returns a new vRA Project service
func NewProjectClient(cfg clients.Config) project.ClientService {
	bearerToken, err := login.GetBearerToken(cfg)
	if err != nil {
		fmt.Printf("Could not get bearer token: %v\n", err)
	}

	vra := clients.NewClientWithAuthentication(cfg, bearerToken)

	return vra.Project
}

// GenerateGetProjectOptions fetches a specific project
func GenerateGetProjectOptions(projectID string) *project.GetProjectParams {

	var params = project.NewGetProjectParams().WithID(
		projectID,
	)
	return params
}

/*func convertUser(users []*v1alpha1.User) []*models.User {
	convertedUsers := []*models.User{}
	for _, v := range users {
		convertedUsers = append(convertedUsers, (*models.User)(v))
	}
	return convertedUsers

}*/

// GenerateCreateProjectOptions generates project creation options
func GenerateCreateProjectOptions(p *v1alpha1.ProjectParameters) *project.CreateProjectParams {
	//viewers:= ([]*models.User) (p.Viewers)
	//[]*models.User{(*models.User) (p.Administrators)}

	var params = project.NewCreateProjectParams().WithBody(
		&models.IaaSProjectSpecification{
			Administrators:               *(*[]*models.User)(unsafe.Pointer(&p.Administrators)),
			Constraints:                  map[string][]models.Constraint{},
			CustomProperties:             map[string]string{},
			Description:                  p.Description,
			MachineNamingTemplate:        p.MachineNamingTemplate,
			Members:                      *(*[]*models.User)(unsafe.Pointer(&p.Members)),
			Name:                         p.Name,
			OperationTimeout:             p.OperationTimeout,
			PlacementPolicy:              p.PlacementPolicy,
			SharedResources:              p.SharedResources,
			Viewers:                      *(*[]*models.User)(unsafe.Pointer(&p.Viewers)),
			ZoneAssignmentConfigurations: []*models.ZoneAssignmentSpecification{},
		},
	)

	return params
}

// GenerateDeleteProjectOptions test
func GenerateDeleteProjectOptions(projectID string) *project.DeleteProjectParams {
	var params = project.NewDeleteProjectParams().WithID(
		projectID,
	)
	return params
}

// GenerateDeleteProjectOptions test
func GenerateUpdateProjectOptions(projectID string, p *v1alpha1.ProjectParameters) *project.UpdateProjectParams {
	var params = project.NewUpdateProjectParams().WithID(
		projectID,
	).WithBody(&models.IaaSProjectSpecification{
		Administrators:               *(*[]*models.User)(unsafe.Pointer(&p.Administrators)),
		Constraints:                  map[string][]models.Constraint{},
		CustomProperties:             map[string]string{},
		Description:                  p.Description,
		MachineNamingTemplate:        p.MachineNamingTemplate,
		Members:                      *(*[]*models.User)(unsafe.Pointer(&p.Members)),
		Name:                         p.Name,
		OperationTimeout:             p.OperationTimeout,
		PlacementPolicy:              p.PlacementPolicy,
		SharedResources:              p.SharedResources,
		Viewers:                      *(*[]*models.User)(unsafe.Pointer(&p.Viewers)),
		ZoneAssignmentConfigurations: []*models.ZoneAssignmentSpecification{},
	})
	return params
}

// GenerateProjectObservation is used to produce v1alpha1.ProjectObservation
func GenerateProjectObservation(project *project.GetProjectOK) v1alpha1.ProjectObservation { // nolint:gocyclo
	if project.Payload == nil {
		return v1alpha1.ProjectObservation{}
	}

	o := v1alpha1.ProjectObservation{
		Name:                  &project.Payload.Name,
		ID:                    *project.Payload.ID,
		Administrators:        *(*[]*v1alpha1.User)(unsafe.Pointer(&project.Payload.Administrators)),
		Members:               *(*[]*v1alpha1.User)(unsafe.Pointer(&project.Payload.Members)),
		Viewers:               *(*[]*v1alpha1.User)(unsafe.Pointer(&project.Payload.Viewers)),
		PlacementPolicy:       project.Payload.PlacementPolicy,
		SharedResources:       project.Payload.SharedResources,
		OperationTimeout:      &project.Payload.OperationTimeout,
		MachineNamingTemplate: project.Payload.MachineNamingTemplate,
		Description:           project.Payload.Description,
	}

	return o
}

func IsResourceUpToDate(desired *v1alpha1.ProjectParameters, current *models.IaaSProject) bool {
	// DisableApiTermination

	// check admins, members, viewers
	if len(current.Administrators) != len(desired.Administrators) || len(current.Members) != len(desired.Members) ||
		len(current.Viewers) != len(desired.Viewers) || current.Name != *desired.Name || current.Description != desired.Description ||
		current.PlacementPolicy != desired.PlacementPolicy || current.MachineNamingTemplate != desired.MachineNamingTemplate ||
		current.SharedResources != desired.SharedResources || current.OperationTimeout != *desired.OperationTimeout {
		return false
	}

	for i := 0; i < len(current.Administrators); i++ {

		if *current.Administrators[i].Email != *desired.Administrators[i].Email || current.Administrators[i].Type != desired.Administrators[i].Type {
			// any admin has changed
			return false
		}
	}
	for i := 0; i < len(current.Members); i++ {

		if *current.Members[i].Email != *desired.Members[i].Email || current.Members[i].Type != desired.Members[i].Type {
			// any member has changed
			return false
		}
	}
	for i := 0; i < len(current.Viewers); i++ {
		if *current.Viewers[i].Email != *desired.Viewers[i].Email || current.Viewers[i].Type != desired.Viewers[i].Type {
			// any viewer has changed
			return false
		}
	}

	return true
}

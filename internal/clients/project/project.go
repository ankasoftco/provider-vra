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
	//fmt.Println(spec)
	converted := *(*[]*models.User)(unsafe.Pointer(&p.Administrators))
	converted2 := *(*[]*models.User)(unsafe.Pointer(&p.Viewers))
	converted3 := *(*[]*models.User)(unsafe.Pointer(&p.Members))
	var params = project.NewCreateProjectParams().WithBody(
		&models.IaaSProjectSpecification{
			Administrators:               converted,
			Constraints:                  map[string][]models.Constraint{},
			CustomProperties:             map[string]string{},
			Description:                  p.Description,
			MachineNamingTemplate:        p.MachineNamingTemplate,
			Members:                      converted3,
			Name:                         p.Name,
			OperationTimeout:             p.OperationTimeout,
			PlacementPolicy:              p.PlacementPolicy,
			SharedResources:              p.SharedResources,
			Viewers:                      converted2,
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

// GenerateProjectObservation is used to produce v1alpha1.ProjectObservation
func GenerateProjectObservation(project *project.GetProjectOK) v1alpha1.ProjectObservation { // nolint:gocyclo
	if project.Payload == nil {
		return v1alpha1.ProjectObservation{}
	}

	o := v1alpha1.ProjectObservation{
		Name:      &project.Payload.Name,
		ProjectID: project.Payload.ID,
	}

	return o
}

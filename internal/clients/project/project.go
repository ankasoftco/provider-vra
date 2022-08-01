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

	var params = project.NewCreateProjectParams().WithBody(
		&models.IaaSProjectSpecification{
			Administrators:               *(*[]*models.User)(unsafe.Pointer(&p.Administrators)),
			Constraints:                  *(*map[string][]models.Constraint)(unsafe.Pointer(&p.Constraints)),
			CustomProperties:             p.CustomProperties,
			Description:                  p.Description,
			MachineNamingTemplate:        p.MachineNamingTemplate,
			Members:                      *(*[]*models.User)(unsafe.Pointer(&p.Members)),
			Name:                         p.Name,
			OperationTimeout:             p.OperationTimeout,
			PlacementPolicy:              p.PlacementPolicy,
			SharedResources:              p.SharedResources,
			Viewers:                      *(*[]*models.User)(unsafe.Pointer(&p.Viewers)),
			ZoneAssignmentConfigurations: *(*[]*models.ZoneAssignmentSpecification)(unsafe.Pointer(&p.ZoneAssignmentConfigurations)),
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

// GenerateUpdateProjectOptions test
func GenerateUpdateProjectOptions(projectID string, p *v1alpha1.ProjectParameters) *project.UpdateProjectParams {
	fmt.Println("UPDATE CALISTIIIII....", p.SharedResources)
	var params = project.NewUpdateProjectParams().WithID(
		projectID,
	).WithBody(&models.IaaSProjectSpecification{
		Administrators:               *(*[]*models.User)(unsafe.Pointer(&p.Administrators)),
		Constraints:                  *(*map[string][]models.Constraint)(unsafe.Pointer(&p.Constraints)),
		CustomProperties:             p.CustomProperties,
		Description:                  p.Description,
		MachineNamingTemplate:        p.MachineNamingTemplate,
		Members:                      *(*[]*models.User)(unsafe.Pointer(&p.Members)),
		Name:                         p.Name,
		OperationTimeout:             p.OperationTimeout,
		PlacementPolicy:              p.PlacementPolicy,
		SharedResources:              p.SharedResources,
		Viewers:                      *(*[]*models.User)(unsafe.Pointer(&p.Viewers)),
		ZoneAssignmentConfigurations: *(*[]*models.ZoneAssignmentSpecification)(unsafe.Pointer(&p.ZoneAssignmentConfigurations)),
	})
	return params
}

// GenerateProjectObservation is used to produce v1alpha1.ProjectObservation
func GenerateProjectObservation(project *project.GetProjectOK) v1alpha1.ProjectObservation { // nolint:gocyclo
	if project.Payload == nil {
		return v1alpha1.ProjectObservation{}
	}

	o := v1alpha1.ProjectObservation{
		Name:                         &project.Payload.Name,
		ID:                           *project.Payload.ID,
		Administrators:               *(*[]*v1alpha1.User)(unsafe.Pointer(&project.Payload.Administrators)),
		Members:                      *(*[]*v1alpha1.User)(unsafe.Pointer(&project.Payload.Members)),
		Viewers:                      *(*[]*v1alpha1.User)(unsafe.Pointer(&project.Payload.Viewers)),
		PlacementPolicy:              project.Payload.PlacementPolicy,
		SharedResources:              project.Payload.SharedResources,
		OperationTimeout:             &project.Payload.OperationTimeout,
		MachineNamingTemplate:        project.Payload.MachineNamingTemplate,
		Description:                  project.Payload.Description,
		ZoneAssignmentConfigurations: *(*[]*v1alpha1.ZoneAssignmentSpecification)(unsafe.Pointer(&project.Payload.Zones)),
		Constraints:                  *(*map[string][]v1alpha1.Constraint)(unsafe.Pointer(&project.Payload.Constraints)),
		CustomProperties:             project.Payload.CustomProperties,
	}

	return o
}

func IsResourceUpToDate(desired *v1alpha1.ProjectParameters, current *models.IaaSProject) bool {

	if current.Name != *desired.Name || current.Description != desired.Description ||
		current.PlacementPolicy != desired.PlacementPolicy || current.MachineNamingTemplate != desired.MachineNamingTemplate ||
		current.OperationTimeout != *desired.OperationTimeout {
		/* fmt.Println("OTHERS FARKLI")
		fmt.Println(current.Name, *desired.Name)
		fmt.Println(current.Description, desired.Description)
		fmt.Println(current.PlacementPolicy, desired.PlacementPolicy)
		fmt.Println(current.MachineNamingTemplate, desired.MachineNamingTemplate)
		fmt.Println(current.OperationTimeout, *desired.OperationTimeout) */
		return false
	}
	if !CompareUsers(current.Administrators, desired.Administrators) {
		fmt.Println("Administrators FARKLI")
		return false
	}
	if !CompareUsers(current.Members, desired.Members) {
		fmt.Println("Members FARKLI")
		return false
	}
	if !CompareUsers(current.Viewers, desired.Viewers) {
		fmt.Println("Viewers FARKLI")
		return false
	}
	if !CompareZones(current.Zones, desired.ZoneAssignmentConfigurations) {
		fmt.Println("Zones FARKLI")
		return false
	}
	if !CompareCustomProperties(current.CustomProperties, desired.CustomProperties) {
		fmt.Println("CUSTOMPROPERTIES FARKLI")
		fmt.Println("CURRENT:", current.CustomProperties)
		fmt.Println("DESIRED:", desired.CustomProperties)
		return false
	}
	if !CompareConstraints(current.Constraints, desired.Constraints) {
		fmt.Println("Constraints FARKLI")
		fmt.Println("CURRENT:", current.Constraints)
		fmt.Println("DESIRED:", desired.Constraints)
		return false
	}

	return true
}

func CompareConstraints(c1 map[string][]models.Constraint, c2 map[string][]v1alpha1.Constraint) bool {
	if len(c1) != len(c2) {
		return false
	}
	for k, v := range c1 {
		if len(v) != len(c2[k]) {
			return false
		}
		for i := 0; i < len(v); i++ {
			if v[i].Expression != c2[k][i].Expression || v[i].Mandatory != c2[k][i].Mandatory {
				return false
			}
		}
	}
	return true

}

func CompareCustomProperties(c1, c2 map[string]string) bool {
	if len(c1) != len(c2) {
		return false
	}
	for k, v := range c1 {
		if v != c2[k] {
			return false
		}
	}
	return true

}

// CompareUsers Comparition of the users to catch differences if exist. (current, desired)
func CompareUsers(users []*models.User, v1alpha1Users []*v1alpha1.User) bool {
	if len(users) != len(v1alpha1Users) {
		return false
	}
	for i := 0; i < len(users); i++ {
		if *users[i].Email != *v1alpha1Users[i].Email || users[i].Type != v1alpha1Users[i].Type {
			return false
		}
	}
	return true
}

// CompareZones Comparition of the zones to catch differences if exist. (current, desired)
func CompareZones(zones []*models.ZoneAssignment, v1aplha1Zones []*v1alpha1.ZoneAssignmentSpecification) bool {

	if len(zones) != len(v1aplha1Zones) {
		return false
	}

	for i, v := range zones {
		if v.CPULimit != v1aplha1Zones[i].CPULimit || v.MaxNumberInstances != v1aplha1Zones[i].MaxNumberInstances ||
			v.MemoryLimitMB != v1aplha1Zones[i].MemoryLimitMB || v.Priority != v1aplha1Zones[i].Priority ||
			v.StorageLimitGB != v1aplha1Zones[i].StorageLimitGB || v.ZoneID != v1aplha1Zones[i].ZoneID {
			return false
		}
	}
	return true
}

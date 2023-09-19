/*
Copyright 2022 The ANKA SOFTWARE Authors.
*/

package blueprint

import (
	"fmt"

	"github.com/crossplane/provider-vraprovider/internal/clients"
	"github.com/crossplane/provider-vraprovider/internal/clients/login"

	"github.com/vmware/vra-sdk-go/pkg/client/blueprint"
)

// NewBlueprintClient returns a new vRA Blueprint service
func NewBlueprintClient(cfg clients.Config) blueprint.ClientService {
	bearerToken, err := login.GetBearerToken(cfg)
	if err != nil {
		fmt.Printf("Could not get bearer token: %v\n", err)
	}

	vra := clients.NewClientWithAuthentication(cfg, bearerToken)

	return vra.Blueprint
}

// GenerateGetBlueprintOptions fetches a specific blueprint
/*
func GenerateGetBlueprintOptions(blueprintID string) *blueprint.GetBlueprintParams {

	var params = blueprint.NewGetBlueprintParams().WithID(
		blueprintID,
	)
	return params
}
*/

/*func convertUser(users []*v1alpha1.User) []*models.User {
	convertedUsers := []*models.User{}
	for _, v := range users {
		convertedUsers = append(convertedUsers, (*models.User)(v))
	}
	return convertedUsers

}*/

// GenerateCreateBlueprintOptions generates blueprint creation options
/*
func GenerateCreateBlueprintOptions(p *v1alpha1.BlueprintParameters) *blueprint.CreateBlueprintParams {

	var params = blueprint.NewCreateBlueprintParams().WithBody(
		&models.IaaSBlueprintSpecification{
			Administrators:        *(*[]*models.User)(unsafe.Pointer(&p.Administrators)),
			Constraints:           *(*map[string][]models.Constraint)(unsafe.Pointer(&p.Constraints)),
			CustomProperties:      p.CustomProperties,
			Description:           p.Description,
			MachineNamingTemplate: p.MachineNamingTemplate,
			Members:               *(*[]*models.User)(unsafe.Pointer(&p.Members)),
			Name:                  p.Name,
			OperationTimeout:      p.OperationTimeout,
			PlacementPolicy:       p.PlacementPolicy,
			//SharedResources:              p.SharedResources,
			Viewers:                      *(*[]*models.User)(unsafe.Pointer(&p.Viewers)),
			ZoneAssignmentConfigurations: *(*[]*models.ZoneAssignmentSpecification)(unsafe.Pointer(&p.ZoneAssignmentConfigurations)),
		},
	)

	return params
}
*/
/*
// GenerateDeleteBlueprintOptions test
func GenerateDeleteBlueprintOptions(blueprintID string) *blueprint.DeleteBlueprintParams {
	var params = blueprint.NewDeleteBlueprintParams().WithID(
		blueprintID,
	)
	return params
}

// GenerateUpdateBlueprintOptions test
func GenerateUpdateBlueprintOptions(blueprintID string, p *v1alpha1.BlueprintParameters) *blueprint.UpdateBlueprintParams {
	//fmt.Println("UPDATE CALISTIIIII....", p.SharedResources)
	var params = blueprint.NewUpdateBlueprintParams().WithID(
		blueprintID,
	).WithBody(&models.IaaSBlueprintSpecification{
		Administrators:        *(*[]*models.User)(unsafe.Pointer(&p.Administrators)),
		Constraints:           *(*map[string][]models.Constraint)(unsafe.Pointer(&p.Constraints)),
		CustomProperties:      p.CustomProperties,
		Description:           p.Description,
		MachineNamingTemplate: p.MachineNamingTemplate,
		Members:               *(*[]*models.User)(unsafe.Pointer(&p.Members)),
		Name:                  p.Name,
		OperationTimeout:      p.OperationTimeout,
		PlacementPolicy:       p.PlacementPolicy,
		//SharedResources:              p.SharedResources,
		Viewers:                      *(*[]*models.User)(unsafe.Pointer(&p.Viewers)),
		ZoneAssignmentConfigurations: *(*[]*models.ZoneAssignmentSpecification)(unsafe.Pointer(&p.ZoneAssignmentConfigurations)),
	})
	return params
}

// GenerateBlueprintObservation is used to produce v1alpha1.BlueprintObservation
func GenerateBlueprintObservation(blueprint *blueprint.GetBlueprintOK) v1alpha1.BlueprintObservation { // nolint:gocyclo
	if blueprint.Payload == nil {
		return v1alpha1.BlueprintObservation{}
	}

	o := v1alpha1.BlueprintObservation{
		Name:            &blueprint.Payload.Name,
		ID:              *blueprint.Payload.ID,
		Administrators:  *(*[]*v1alpha1.User)(unsafe.Pointer(&blueprint.Payload.Administrators)),
		Members:         *(*[]*v1alpha1.User)(unsafe.Pointer(&blueprint.Payload.Members)),
		Viewers:         *(*[]*v1alpha1.User)(unsafe.Pointer(&blueprint.Payload.Viewers)),
		PlacementPolicy: blueprint.Payload.PlacementPolicy,
		//SharedResources:              blueprint.Payload.SharedResources,
		OperationTimeout:             &blueprint.Payload.OperationTimeout,
		MachineNamingTemplate:        blueprint.Payload.MachineNamingTemplate,
		Description:                  blueprint.Payload.Description,
		ZoneAssignmentConfigurations: *(*[]*v1alpha1.ZoneAssignmentSpecification)(unsafe.Pointer(&blueprint.Payload.Zones)),
		Constraints:                  *(*map[string][]v1alpha1.Constraint)(unsafe.Pointer(&blueprint.Payload.Constraints)),
		CustomProperties:             blueprint.Payload.CustomProperties,
	}

	return o
}

func IsResourceUpToDate(desired *v1alpha1.BlueprintParameters, current *models.IaaSBlueprint) bool {

	if current.Name != *desired.Name || current.Description != desired.Description ||
		current.PlacementPolicy != desired.PlacementPolicy || current.MachineNamingTemplate != desired.MachineNamingTemplate ||
		current.OperationTimeout != *desired.OperationTimeout {
		return false
	}
	if !CompareUsers(current.Administrators, desired.Administrators) {
		return false
	}
	if !CompareUsers(current.Members, desired.Members) {
		return false
	}
	if !CompareUsers(current.Viewers, desired.Viewers) {
		return false
	}
	if !CompareZones(current.Zones, desired.ZoneAssignmentConfigurations) {
		return false
	}
	if !CompareCustomProperties(current.CustomProperties, desired.CustomProperties) {
		return false
	}
	if !CompareConstraints(current.Constraints, desired.Constraints) {
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
			if *v[i].Expression != *c2[k][i].Expression || *v[i].Mandatory != *c2[k][i].Mandatory {
				fmt.Println("EXPRESSION FARKLIYDI:", *v[i].Expression, *c2[k][i].Expression)
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
*/

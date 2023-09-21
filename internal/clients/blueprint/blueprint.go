/*
Copyright 2022 The ANKA SOFTWARE Authors.
*/

package blueprint

import (
	"fmt"

	"github.com/crossplane/provider-vraprovider/apis/vra/v1alpha1"
	"github.com/crossplane/provider-vraprovider/internal/clients"
	"github.com/crossplane/provider-vraprovider/internal/clients/login"
	"github.com/go-openapi/strfmt"

	"github.com/vmware/vra-sdk-go/pkg/client/blueprint"
	"github.com/vmware/vra-sdk-go/pkg/models"
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

// GenerateGetBlueprintOptions fetches a specific blueprintGetBlueprintParams
func GenerateGetBlueprintOptions(blueprintID string) *blueprint.GetBlueprintUsingGET1Params {
	var params = blueprint.NewGetBlueprintUsingGET1Params().WithBlueprintID(
		strfmt.UUID(blueprintID),
	)
	return params
}

// GenerateCreateBlueprintOptions generates blueprint creation options
func GenerateCreateBlueprintOptions(b *v1alpha1.BlueprintParameters) *blueprint.CreateBlueprintUsingPOST1Params {

	var params = blueprint.NewCreateBlueprintUsingPOST1Params().WithBlueprint(
		&models.Blueprint{
			Content:     b.Content,
			Description: b.Description,
			Name:        b.Name,
			ProjectID:   b.ProjectID,
		},
	)
	return params
}

// GenerateDeleteBlueprintOptions test
func GenerateDeleteBlueprintOptions(blueprintID string) *blueprint.DeleteBlueprintUsingDELETE1Params {
	var params = blueprint.NewDeleteBlueprintUsingDELETE1Params().WithBlueprintID(
		strfmt.UUID(blueprintID),
	)
	return params
}

// GenerateUpdateBlueprintOptions test
func GenerateUpdateBlueprintOptions(blueprintID string, b *v1alpha1.BlueprintParameters) *blueprint.UpdateBlueprintUsingPUT1Params {
	var params = blueprint.NewUpdateBlueprintUsingPUT1Params().WithBlueprintID(
		strfmt.UUID(blueprintID),
	).WithBlueprint(&models.Blueprint{
		Content:           b.Content,
		Description:       b.Description,
		Name:              b.Name,
		ProjectID:         b.ProjectID,
		CreatedBy:         b.CreatedBy,
		ContentSourceType: b.ContentSourceType,
	})
	return params
}

// GenerateBlueprintObservation is used to produce v1alpha1.BlueprintObservation
func GenerateBlueprintObservation(b *blueprint.GetBlueprintUsingGET1OK) v1alpha1.BlueprintObservation { // nolint:gocyclo
	if b.Payload == nil {
		return v1alpha1.BlueprintObservation{}
	}

	o := v1alpha1.BlueprintObservation{
		Content:           b.Payload.Content,
		Description:       b.Payload.Description,
		Name:              b.Payload.Name,
		ProjectID:         b.Payload.ProjectID,
		CreatedBy:         b.Payload.CreatedBy,
		ContentSourceType: b.Payload.ContentSourceType,
	}

	return o
}

func IsResourceUpToDate(desired *v1alpha1.BlueprintParameters, current *models.Blueprint) bool {
	if current.Name != desired.Name || current.Description != desired.Description || current.Content != desired.Content {
		return false
	}
	return true
}

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

package blueprint

import (
	"github.com/crossplane/provider-vra/apis/blueprint/v1alpha1"
	"github.com/crossplane/provider-vra/internal/clients"
	"github.com/go-openapi/strfmt"
	"github.com/vmware/vra-sdk-go/pkg/models"

	"github.com/vmware/vra-sdk-go/pkg/client/blueprint"
)

// NewBlueprintClient returns a new vRA Blueprint service
func NewBlueprintClient(cfg clients.Config) blueprint.ClientService {
	vra := clients.NewClient2(cfg)

	return vra.Blueprint
}

// GenerateCreateBLueprintOptions generates deployment creation options
func GenerateCreateBlueprintOptions(p *v1alpha1.BlueprintParameters) *blueprint.CreateBlueprintUsingPOST1Params {
	return blueprint.NewCreateBlueprintUsingPOST1Params().WithBlueprint(&models.Blueprint{
		Name:      p.Name,
		ProjectID: p.ProjectID,
		Content:   "format",
	})
}

func GenerateGetBlueprintOptions(blueprintID string) *blueprint.GetBlueprintUsingGET1Params {
	return blueprint.NewGetBlueprintUsingGET1Params().WithBlueprintID(strfmt.UUID(blueprintID))
}

func GenerateDeleteBlueprintOptions(blueprintID string) *blueprint.DeleteBlueprintUsingDELETE1Params {
	return blueprint.NewDeleteBlueprintUsingDELETE1Params().WithBlueprintID(strfmt.UUID(blueprintID))
}

func GenerateObservation(blueprint *blueprint.GetBlueprintUsingGET1OK) v1alpha1.BlueprintObservation { // nolint:gocyclo
	if blueprint.Payload == nil {
		return v1alpha1.BlueprintObservation{}
	}

	o := v1alpha1.BlueprintObservation{
		Name:        &blueprint.Payload.Name,
		BlueprintID: strfmt.UUID(blueprint.Payload.ID),
	}

	return o
}

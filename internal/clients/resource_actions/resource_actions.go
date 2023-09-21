/*
Copyright 2022 The ANKA SOFTWARE Authors.
*/

package deployment

import (
	"fmt"

	"github.com/go-openapi/strfmt"
	"github.com/vmware/vra-sdk-go/pkg/client/resource_actions"
	"github.com/vmware/vra-sdk-go/pkg/models"

	"github.com/crossplane/provider-vraprovider/apis/vra/v1alpha1"
	"github.com/crossplane/provider-vraprovider/internal/clients"
	"github.com/crossplane/provider-vraprovider/internal/clients/login"
)

// ResourceActionsClient returns a new vRA Deployment service
func ResourceActionsClient(cfg clients.Config) resource_actions.ClientService {
	bearerToken, err := login.GetBearerToken(cfg)
	if err != nil {
		fmt.Printf("Could not get bearer token: %v\n", err)
	}

	vra := clients.NewClientWithAuthentication(cfg, bearerToken)

	return vra.ResourceActions
}

// GenerateGetDeploymentOptions fetches a specific deployment
func GenerateGetResourceActionsOptions(deployment_actionID string, d *v1alpha1.ResourceActionsParameters) *resource_actions.GetResourceActionUsingGET5Params {
	var params = resource_actions.NewGetResourceActionUsingGET5Params().WithResourceID(
		strfmt.UUID(d.ResourceID),
	).WithActionID(
		deployment_actionID,
	)

	return params
}

// GenerateDeploymentObservation is used to produce v1alpha1.DeploymentObservation
func GenerateResourceActionsObservation(d *resource_actions.GetResourceActionUsingGET5OK) v1alpha1.ResourceActionsObservation { // nolint:gocyclo
	if d.Payload == nil {
		return v1alpha1.ResourceActionsObservation{}
	}

	o := v1alpha1.ResourceActionsObservation{
		Name: &d.Payload.DisplayName,
		ID:   d.Payload.ID,
	}

	return o
}

// GenerateCreateProjectOptions generates project creation options
func GenerateCreateResourceActionsOptions(d *v1alpha1.ResourceActionsParameters) *resource_actions.SubmitResourceActionRequestUsingPOST5Params {

	var params = resource_actions.NewSubmitResourceActionRequestUsingPOST5Params().WithResourceID(
		strfmt.UUID(d.ResourceID),
	).WithActionRequest(
		&models.ResourceActionRequest{
			ActionID: d.ActionID,
			Inputs:   d.Inputs,
			Reason:   d.Reason,
		},
	)

	return params
}

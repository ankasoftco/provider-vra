/*
Copyright 2022 The ANKA SOFTWARE Authors.
*/

package deployment

import (
	"fmt"

	"github.com/go-openapi/strfmt"
	"github.com/vmware/vra-sdk-go/pkg/client/deployment_actions"
	"github.com/vmware/vra-sdk-go/pkg/models"

	"github.com/crossplane/provider-vraprovider/apis/vra/v1alpha1"
	"github.com/crossplane/provider-vraprovider/internal/clients"
	"github.com/crossplane/provider-vraprovider/internal/clients/login"
)

// DeploymentActionsClient returns a new vRA Deployment service
func DeploymentActionsClient(cfg clients.Config) deployment_actions.ClientService {
	bearerToken, err := login.GetBearerToken(cfg)
	if err != nil {
		fmt.Printf("Could not get bearer token: %v\n", err)
	}

	vra := clients.NewClientWithAuthentication(cfg, bearerToken)

	return vra.DeploymentActions
}

// GenerateGetDeploymentOptions fetches a specific deployment
func GenerateGetDeploymentActionsOptions(deployment_actionID string, d *v1alpha1.DeploymentActionsParameters) *deployment_actions.GetDeploymentActionUsingGET2Params {
	var params = deployment_actions.NewGetDeploymentActionUsingGET2Params().WithDeploymentID(
		strfmt.UUID(d.DeploymentID),
	).WithActionID(
		deployment_actionID,
	)

	return params
}

// GenerateDeploymentObservation is used to produce v1alpha1.DeploymentObservation
func GenerateDeploymentActionsObservation(d *deployment_actions.GetDeploymentActionUsingGET2OK) v1alpha1.DeploymentActionsObservation { // nolint:gocyclo
	if d.Payload == nil {
		return v1alpha1.DeploymentActionsObservation{}
	}

	o := v1alpha1.DeploymentActionsObservation{
		Name: &d.Payload.DisplayName,
		ID:   d.Payload.ID,
	}

	return o
}

// GenerateCreateProjectOptions generates project creation options
func GenerateCreateDeploymentActionsOptions(d *v1alpha1.DeploymentActionsParameters) *deployment_actions.SubmitDeploymentActionRequestUsingPOST2Params {

	var params = deployment_actions.NewSubmitDeploymentActionRequestUsingPOST2Params().WithDeploymentID(
		strfmt.UUID(d.DeploymentID),
	).WithActionRequest(
		&models.ResourceActionRequest{
			ActionID: d.ActionID,
			Inputs:   d.Inputs,
			Reason:   d.Reason,
		},
	)

	return params
}

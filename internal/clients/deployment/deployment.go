/*
Copyright 2022 The ANKA SOFTWARE Authors.
*/

package deployment

import (
	"fmt"

	"github.com/go-openapi/strfmt"

	"github.com/vmware/vra-sdk-go/pkg/client/deployments"

	"github.com/crossplane/provider-vra/apis/deployment/v1alpha1"
	"github.com/crossplane/provider-vra/internal/clients"
	"github.com/crossplane/provider-vra/internal/clients/login"
)

// NewDeploymentClient returns a new vRA Deployment service
func NewDeploymentClient(cfg clients.Config) deployments.ClientService {
	bearerToken, err := login.GetBearerToken(cfg)
	if err != nil {
		fmt.Printf("Could not get bearer token: %v\n", err)
	}

	vra := clients.NewClientWithAuthentication(cfg, bearerToken)

	return vra.Deployments
}

// GenerateGetDeploymentOptions fetches a specific deployment
func GenerateGetDeploymentOptions(deploymentID string) *deployments.GetDeploymentByIDV3UsingGETParams {
	var params = deployments.NewGetDeploymentByIDV3UsingGETParams().WithDeploymentID(
		strfmt.UUID(deploymentID),
	)

	return params
}

// GenerateDeleteDeploymentOptions a
func GenerateDeleteDeploymentOptions(deploymentID string) *deployments.DeleteDeploymentUsingDELETE2Params {
	var params = deployments.NewDeleteDeploymentUsingDELETE2Params().WithDeploymentID(
		strfmt.UUID(deploymentID),
	)

	return params
}

// GenerateDeploymentObservation is used to produce v1alpha1.DeploymentObservation
func GenerateDeploymentObservation(d *deployments.GetDeploymentByIDV3UsingGETOK) v1alpha1.DeploymentObservation { // nolint:gocyclo
	if d.Payload == nil {
		return v1alpha1.DeploymentObservation{}
	}

	o := v1alpha1.DeploymentObservation{
		Name:         *d.Payload.Name,
		DeploymentID: d.Payload.ID,
		Status:       d.Payload.Status,
	}

	return o
}

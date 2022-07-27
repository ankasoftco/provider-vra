/*
Copyright 2022 The ANKA SOFTWARE Authors.
*/

package catalogitem

import (
	"fmt"

	"github.com/vmware/vra-sdk-go/pkg/client/catalog_items"
	"github.com/vmware/vra-sdk-go/pkg/models"

	"github.com/crossplane/provider-vra/apis/deployment/v1alpha1"
	"github.com/crossplane/provider-vra/internal/clients"
	"github.com/crossplane/provider-vra/internal/clients/login"
)

// NewCatalogItemClient returns a new vRA Catalog Item service
func NewCatalogItemClient(cfg clients.Config) catalog_items.ClientService {
	bearerToken, err := login.GetBearerToken(cfg)
	if err != nil {
		fmt.Printf("Could not get bearer token: %v\n", err)
	}

	vra := clients.NewClientWithAuthentication(cfg, bearerToken)

	return vra.CatalogItems
}

// GenerateRequestCatalogItemOptions creates deployments from a catalog item.
func GenerateRequestCatalogItemOptions(d *v1alpha1.DeploymentParameters) *catalog_items.RequestCatalogItemInstancesUsingPOST1Params {
	var params = catalog_items.NewRequestCatalogItemInstancesUsingPOST1Params().WithID(
		d.CatalogItemID,
	).WithRequest(
		&models.CatalogItemRequest{
			DeploymentName: d.DeploymentName,
			Inputs:         d.Inputs,
			ProjectID:      d.ProjectID,
		},
	)

	if d.Reason != "" {
		params.Request.Reason = d.Reason
	}

	if d.Version != "" {
		params.Request.Version = d.Version
	}

	if d.BulkRequestCount != nil {
		params.Request.BulkRequestCount = d.BulkRequestCount
	}

	return params
}

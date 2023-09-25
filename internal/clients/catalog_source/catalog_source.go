/*
Copyright 2022 The ANKA SOFTWARE Authors.
*/

package catalog_sources

import (
	"fmt"

	"github.com/crossplane/provider-vraprovider/apis/vra/v1alpha1"
	"github.com/crossplane/provider-vraprovider/internal/clients"
	"github.com/crossplane/provider-vraprovider/internal/clients/login"
	"github.com/go-openapi/strfmt"

	"github.com/vmware/vra-sdk-go/pkg/client/catalog_sources"
	"github.com/vmware/vra-sdk-go/pkg/models"
)

// NewCatalogSourceClient returns a new vRA CatalogSource service
func NewCatalogSourceClient(cfg clients.Config) catalog_sources.ClientService {
	bearerToken, err := login.GetBearerToken(cfg)
	if err != nil {
		fmt.Printf("Could not get bearer token: %v\n", err)
	}

	vra := clients.NewClientWithAuthentication(cfg, bearerToken)

	return vra.CatalogSources
}

// GenerateGetCatalogSourceOptions fetches a specific catalog_sourcesGetCatalogSourceParams
func GenerateGetCatalogSourceOptions(catalog_sourcesID string) *catalog_sources.GetUsingGET2Params {
	var params = catalog_sources.NewGetUsingGET2Params().WithSourceID(
		strfmt.UUID(catalog_sourcesID),
	)
	return params
}

// GenerateCreateCatalogSourceOptions generates catalog_sources creation options
func GenerateCreateCatalogSourceOptions(b *v1alpha1.CatalogSourceParameters) *catalog_sources.PostUsingPOST2Params {

	var params = catalog_sources.NewPostUsingPOST2Params().WithSource(
		&models.CatalogSource{
			Config:      b.Config,
			Description: b.Description,
			Global:      b.Global,
			Name:        b.Name,
			ProjectID:   b.ProjectID,
			TypeID:      b.TypeID,
		},
	)
	return params
}

// GenerateDeleteCatalogSourceOptions test
func GenerateDeleteCatalogSourceOptions(catalog_sourcesID string) *catalog_sources.DeleteUsingDELETE4Params {
	var params = catalog_sources.NewDeleteUsingDELETE4Params().WithSourceID(
		strfmt.UUID(catalog_sourcesID),
	)
	return params
}

// GenerateUpdateCatalogSourceOptions test
func GenerateUpdateCatalogSourceOptions(catalog_sourcesID string, b *v1alpha1.CatalogSourceParameters) *catalog_sources.PostUsingPOST2Params {
	var params = catalog_sources.NewPostUsingPOST2Params().WithSource(
		&models.CatalogSource{
			Config:      b.Config,
			Description: b.Description,
			Global:      b.Global,
			Name:        b.Name,
			ProjectID:   b.ProjectID,
			TypeID:      b.TypeID,
		},
	)
	return params
}

// GenerateCatalogSourceObservation is used to produce v1alpha1.CatalogSourceObservation
func GenerateCatalogSourceObservation(b *catalog_sources.GetUsingGET2OK) v1alpha1.CatalogSourceObservation { // nolint:gocyclo
	if b.Payload == nil {
		return v1alpha1.CatalogSourceObservation{}
	}

	o := v1alpha1.CatalogSourceObservation{
		ID:               b.Payload.ID,
		CreatedBy:        b.Payload.CreatedBy,
		Description:      b.Payload.Description,
		Name:             b.Payload.Name,
		ProjectID:        b.Payload.ProjectID,
		TypeID:           b.Payload.TypeID,
		LastImportErrors: b.Payload.LastImportErrors,
		ItemsFound:       b.Payload.ItemsFound,
		ItemsImported:    b.Payload.ItemsImported,
		LastUpdatedBy:    b.Payload.LastUpdatedBy,
		Global:           b.Payload.Global,
		IconID:           b.Payload.IconID,
		Config:           fmt.Sprintf("%v", b.Payload.Config),
	}

	return o
}

func IsResourceUpToDate(desired *v1alpha1.CatalogSourceParameters, current *models.CatalogSource) bool {
	if *current.Name != *desired.Name || current.Description != desired.Description {
		return false
	}
	return true
}

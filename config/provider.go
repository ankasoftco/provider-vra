/*
Copyright 2021 Upbound Inc.
*/

package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	blockDevice "github.com/ankasoftco/upjet-provider-vra/config/block_device"
	blueprint "github.com/ankasoftco/upjet-provider-vra/config/blueprint"
	catalogItemEntitlement "github.com/ankasoftco/upjet-provider-vra/config/catalog_item"
	catalogSource "github.com/ankasoftco/upjet-provider-vra/config/catalog_source"
	contentSource "github.com/ankasoftco/upjet-provider-vra/config/content_source"
	deployment "github.com/ankasoftco/upjet-provider-vra/config/deployment"
	fabric "github.com/ankasoftco/upjet-provider-vra/config/fabric"
	flavorProfile "github.com/ankasoftco/upjet-provider-vra/config/flavor_profile"
	imageProfile "github.com/ankasoftco/upjet-provider-vra/config/image_profile"
	integration "github.com/ankasoftco/upjet-provider-vra/config/integration"
	loadBalancer "github.com/ankasoftco/upjet-provider-vra/config/load_balancer"
	project "github.com/ankasoftco/upjet-provider-vra/config/project"
	storage "github.com/ankasoftco/upjet-provider-vra/config/storage"
	ujconfig "github.com/upbound/upjet/pkg/config"
)

const (
	resourcePrefix = "vra"
	modulePath     = "github.com/ankasoftco/upjet-provider-vra"
)

//go:embed schema.json
var providerSchema string

//go:embed provider-metadata.yaml
var providerMetadata string

// GetProvider returns provider configuration
func GetProvider() *ujconfig.Provider {
	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithRootGroup("crossplane.io"),
		ujconfig.WithIncludeList(ExternalNameConfigured()),
		ujconfig.WithFeaturesPackage("internal/features"),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
		))

	for _, configure := range []func(provider *ujconfig.Provider){
		// add custom config functions
		project.Configure,
		blueprint.Configure,
		deployment.Configure,
		fabric.Configure,
		blockDevice.Configure,
		flavorProfile.Configure,
		imageProfile.Configure,
		storage.Configure,
		catalogSource.Configure,
		catalogItemEntitlement.Configure,
		contentSource.Configure,
		integration.Configure,
		loadBalancer.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}

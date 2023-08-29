/*
Copyright 2021 Upbound Inc.
*/

package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	blockDevice "github.com/ankasoftco/upjet-provider-vra/config/block_device"

	ujconfig "github.com/upbound/upjet/pkg/config"

	blueprint "github.com/ankasoftco/upjet-provider-vra/config/blueprint"
	deployment "github.com/ankasoftco/upjet-provider-vra/config/deployment"
<<<<<<< HEAD
	fabric "github.com/ankasoftco/upjet-provider-vra/config/fabric"
=======
	project "github.com/ankasoftco/upjet-provider-vra/config/project"
>>>>>>> a5278da (block device and block device snapshot added)
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
		blockDeviceSnapshot.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}

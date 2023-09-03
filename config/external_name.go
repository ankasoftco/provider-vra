/*
Copyright 2022 Upbound Inc.
*/

package config

import "github.com/upbound/upjet/pkg/config"

// ExternalNameConfigs contains all external name configurations for this
// provider.
var ExternalNameConfigs = map[string]config.ExternalName{
	// Import requires using a randomly generated ID from provider: nl-2e21sda
	"vra_project":                    config.IdentifierFromProvider,
	"vra_blueprint":                  config.IdentifierFromProvider,
	"vra_blueprint_version":          config.IdentifierFromProvider,
	"vra_cloud_account_aws":          config.IdentifierFromProvider,
	"vra_cloud_account_azure":        config.IdentifierFromProvider,
	"vra_cloud_account_gcp":          config.IdentifierFromProvider,
	"vra_cloud_account_nsxt":         config.IdentifierFromProvider,
	"vra_cloud_account_vmc":          config.IdentifierFromProvider,
	"vra_cloud_account_vsphere":      config.IdentifierFromProvider,
	"vra_deployment":                 config.IdentifierFromProvider,
	"vra_fabric_compute":             config.IdentifierFromProvider,
	"vra_fabric_datastore_vsphere":   config.IdentifierFromProvider,
	"vra_fabric_network_vsphere":     config.IdentifierFromProvider,
	"vra_flavor_profile":             config.IdentifierFromProvider,
	"vra_image_profile":              config.IdentifierFromProvider,
	"vra_storage_profile":            config.IdentifierFromProvider,
	"vra_storage_profile_aws":        config.IdentifierFromProvider,
	"vra_storage_profile_azure":      config.IdentifierFromProvider,
	"vra_storage_profile_vsphere":    config.IdentifierFromProvider,
	"vra_block_device":               config.IdentifierFromProvider,
	"vra_block_device_snapshot":      config.IdentifierFromProvider,
	"vra_catalog_source_blueprint":   config.IdentifierFromProvider,
	"vra_catalog_source_entitlement": config.IdentifierFromProvider,
	"vra_catalog_item_entitlement":   config.IdentifierFromProvider,
	"vra_content_source":             config.IdentifierFromProvider,
	"vra_integration":                config.IdentifierFromProvider,
	"vra_load_balancer":              config.IdentifierFromProvider,
	"vra_machine":                    config.IdentifierFromProvider,
	"vra_network":                    config.IdentifierFromProvider,
	"vra_network_ip_range":           config.IdentifierFromProvider,
	"vra_network_profile":            config.IdentifierFromProvider,
	"vra_zone":                       config.IdentifierFromProvider,
}

// ExternalNameConfigurations applies all external name configs listed in the
// table ExternalNameConfigs and sets the version of those resources to v1beta1
// assuming they will be tested.
func ExternalNameConfigurations() config.ResourceOption {
	return func(r *config.Resource) {
		if e, ok := ExternalNameConfigs[r.Name]; ok {
			r.ExternalName = e
		}
	}
}

// ExternalNameConfigured returns the list of all resources whose external name
// is configured manually.
func ExternalNameConfigured() []string {
	l := make([]string, len(ExternalNameConfigs))
	i := 0
	for name := range ExternalNameConfigs {
		// $ is added to match the exact string since the format is regex.
		l[i] = name + "$"
		i++
	}
	return l
}

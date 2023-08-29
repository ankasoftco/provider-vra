package storage

import "github.com/upbound/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("vra_storage_profile", func(r *config.Resource) {

		r.ShortGroup = "storage"
		r.Version = "v1alpha1"
	})

	p.AddResourceConfigurator("vra_storage_profile_aws", func(r *config.Resource) {

		r.ShortGroup = "storage"
		r.Version = "v1alpha1"
	})

	p.AddResourceConfigurator("vra_storage_profile_azure", func(r *config.Resource) {

		r.ShortGroup = "storage"
		r.Version = "v1alpha1"
	})

	p.AddResourceConfigurator("vra_storage_profile_vsphere", func(r *config.Resource) {

		r.ShortGroup = "storage"
		r.Version = "v1alpha1"
	})
}

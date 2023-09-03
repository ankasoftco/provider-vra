package cloudAccount

import "github.com/upbound/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("vra_cloud_account_aws", func(r *config.Resource) {
		r.Version = "v1alpha1"
		r.ShortGroup = "cloudAccount"
	})

	p.AddResourceConfigurator("vra_cloud_account_azure", func(r *config.Resource) {
		r.Version = "v1alpha1"
		r.ShortGroup = "cloudAccount"
	})

	p.AddResourceConfigurator("vra_cloud_account_gcp", func(r *config.Resource) {
		r.Version = "v1alpha1"
		r.ShortGroup = "cloudAccount"
	})

	p.AddResourceConfigurator("vra_cloud_account_nsxt", func(r *config.Resource) {
		r.Version = "v1alpha1"
		r.ShortGroup = "cloudAccount"
	})

	p.AddResourceConfigurator("vra_cloud_account_nsxv", func(r *config.Resource) {
		r.Version = "v1alpha1"
		r.ShortGroup = "cloudAccount"
	})

	p.AddResourceConfigurator("vra_cloud_account_vmc", func(r *config.Resource) {
		r.Version = "v1alpha1"
		r.ShortGroup = "cloudAccount"
	})

	p.AddResourceConfigurator("vra_cloud_account_vsphere", func(r *config.Resource) {
		r.Version = "v1alpha1"
		r.ShortGroup = "cloudAccount"
	})
}

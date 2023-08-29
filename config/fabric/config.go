package fabric

import "github.com/upbound/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("vra_fabric_compute", func(r *config.Resource) {
		r.Version = "v1alpha1"
		r.ShortGroup = "fabric"
	})

	p.AddResourceConfigurator("vra_fabric_datastore_vsphere", func(r *config.Resource) {
		r.Version = "v1alpha1"
		r.ShortGroup = "fabric"
	})

	p.AddResourceConfigurator("vra_fabric_network_vsphere", func(r *config.Resource) {
		r.Version = "v1alpha1"
		r.ShortGroup = "fabric"
	})
}

package flavorprofile

import "github.com/upbound/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("vra_flavor_profile", func(r *config.Resource) {
		r.ShortGroup = "flavorprofile"
		r.Version = "v1alpha1"
	})
}

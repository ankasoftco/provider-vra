package imageProfile

import "github.com/upbound/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("vra_image_profile", func(r *config.Resource) {
		r.ShortGroup = "imageprofile"
		r.Version = "v1alpha1"
	})
}

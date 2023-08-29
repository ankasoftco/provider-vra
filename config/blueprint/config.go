package repository

import "github.com/upbound/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
    p.AddResourceConfigurator("vra_blueprint", func(r *config.Resource) {
  
        r.ShortGroup = "blueprint"     
        r.Version = "v1alpha1"      
    })

    p.AddResourceConfigurator("vra_blueprint_version", func(r *config.Resource) {
  
        r.ShortGroup = "blueprint.version"     
        r.Version = "v1alpha1"      
    })
}
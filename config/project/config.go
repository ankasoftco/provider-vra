package repository

import "github.com/upbound/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
    p.AddResourceConfigurator("vra_project", func(r *config.Resource) {

        r.ShortGroup = "project"     
        r.Version = "v1alpha1"      
    })
}
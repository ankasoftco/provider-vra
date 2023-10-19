package contentsource

import "github.com/upbound/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("vra_content_source", func(r *config.Resource) {
		r.ShortGroup = "contentsource"
		r.Kind = "ContentSource"
		r.Version = "v1alpha1"
		r.References["project_id"] = config.Reference{
			Type: "github.com/ankasoftco/provider-vra/apis/project/v1alpha1.Project",
		}
	})
}

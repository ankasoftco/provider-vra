package catalogSource

import "github.com/upbound/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("vra_catalog_source_blueprint", func(r *config.Resource) {
		r.ShortGroup = "catalogsourceblueprint"
		r.Kind = "CatalogSourceBlueprint"
		r.Version = "v1alpha1"
		r.References["project_id"] = config.Reference{
			Type: "github.com/ankasoftco/provider-vra/apis/project/v1alpha1.Project",
		}
	})

	p.AddResourceConfigurator("vra_catalog_source_entitlement", func(r *config.Resource) {
		r.ShortGroup = "catalogsourceentitlement"
		r.Kind = "CatalogSourceEntitlement"
		r.Version = "v1alpha1"
		r.References["project_id"] = config.Reference{
			Type: "github.com/ankasoftco/provider-vra/apis/project/v1alpha1.Project",
		}
	})
}

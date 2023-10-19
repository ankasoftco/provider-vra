package catalogitem

import "github.com/upbound/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("vra_catalog_item_entitlement", func(r *config.Resource) {
		r.ShortGroup = "catalogitementitlement"
		r.Kind = "CatalogItemEntitlement"
		r.Version = "v1alpha1"
		r.References["project_id"] = config.Reference{
			Type: "github.com/ankasoftco/provider-vra/apis/project/v1alpha1.Project",
		}
	})
}

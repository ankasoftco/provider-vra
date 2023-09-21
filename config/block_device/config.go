package blockDevice

import "github.com/upbound/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("vra_block_device", func(r *config.Resource) {
		r.ShortGroup = "blockdevice"
		r.Kind = "BlockDevice"
		r.Version = "v1alpha1"
		r.References["project_id"] = config.Reference{
			Type: "github.com/ankasoftco/provider-vra/apis/project/v1alpha1.Project",
		}
	})

	p.AddResourceConfigurator("vra_block_device_snapshot", func(r *config.Resource) {
		r.ShortGroup = "blockdevice"
		r.Kind = "BlockDeviceSnapshot"
		r.Version = "v1alpha1"
		r.References["block_device_id"] = config.Reference{
			Type: "BlockDevice",
		}
	})
}

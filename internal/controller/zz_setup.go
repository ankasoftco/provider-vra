/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	blockdevice "github.com/ankasoftco/upjet-provider-vra/internal/controller/blockdevice/blockdevice"
	blockdevicesnapshot "github.com/ankasoftco/upjet-provider-vra/internal/controller/blockdevice/blockdevicesnapshot"
	blueprint "github.com/ankasoftco/upjet-provider-vra/internal/controller/blueprint/blueprint"
	version "github.com/ankasoftco/upjet-provider-vra/internal/controller/blueprint/version"
	catalogitementitlement "github.com/ankasoftco/upjet-provider-vra/internal/controller/catalogitementitlement/catalogitementitlement"
	catalogsourceblueprint "github.com/ankasoftco/upjet-provider-vra/internal/controller/catalogsourceblueprint/catalogsourceblueprint"
	catalogsourceentitlement "github.com/ankasoftco/upjet-provider-vra/internal/controller/catalogsourceentitlement/catalogsourceentitlement"
	contentsource "github.com/ankasoftco/upjet-provider-vra/internal/controller/contentsource/contentsource"
	deployment "github.com/ankasoftco/upjet-provider-vra/internal/controller/deployment/deployment"
	compute "github.com/ankasoftco/upjet-provider-vra/internal/controller/fabric/compute"
	datastorevsphere "github.com/ankasoftco/upjet-provider-vra/internal/controller/fabric/datastorevsphere"
	networkvsphere "github.com/ankasoftco/upjet-provider-vra/internal/controller/fabric/networkvsphere"
	profile "github.com/ankasoftco/upjet-provider-vra/internal/controller/flavorprofile/profile"
	profileimageprofile "github.com/ankasoftco/upjet-provider-vra/internal/controller/imageprofile/profile"
	integration "github.com/ankasoftco/upjet-provider-vra/internal/controller/integration/integration"
	loadbalancer "github.com/ankasoftco/upjet-provider-vra/internal/controller/loadbalancer/loadbalancer"
	machine "github.com/ankasoftco/upjet-provider-vra/internal/controller/machine/machine"
	project "github.com/ankasoftco/upjet-provider-vra/internal/controller/project/project"
	providerconfig "github.com/ankasoftco/upjet-provider-vra/internal/controller/providerconfig"
	profilestorage "github.com/ankasoftco/upjet-provider-vra/internal/controller/storage/profile"
	profileaws "github.com/ankasoftco/upjet-provider-vra/internal/controller/storage/profileaws"
	profileazure "github.com/ankasoftco/upjet-provider-vra/internal/controller/storage/profileazure"
	profilevsphere "github.com/ankasoftco/upjet-provider-vra/internal/controller/storage/profilevsphere"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		blockdevice.Setup,
		blockdevicesnapshot.Setup,
		blueprint.Setup,
		version.Setup,
		catalogitementitlement.Setup,
		catalogsourceblueprint.Setup,
		catalogsourceentitlement.Setup,
		contentsource.Setup,
		deployment.Setup,
		compute.Setup,
		datastorevsphere.Setup,
		networkvsphere.Setup,
		profile.Setup,
		profileimageprofile.Setup,
		integration.Setup,
		loadbalancer.Setup,
		machine.Setup,
		project.Setup,
		providerconfig.Setup,
		profilestorage.Setup,
		profileaws.Setup,
		profileazure.Setup,
		profilevsphere.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

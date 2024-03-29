/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	blockdevice "github.com/ankasoftco/provider-vra/internal/controller/blockdevice/blockdevice"
	blockdevicesnapshot "github.com/ankasoftco/provider-vra/internal/controller/blockdevice/blockdevicesnapshot"
	blueprint "github.com/ankasoftco/provider-vra/internal/controller/blueprint/blueprint"
	version "github.com/ankasoftco/provider-vra/internal/controller/blueprint/version"
	catalogitementitlement "github.com/ankasoftco/provider-vra/internal/controller/catalogitementitlement/catalogitementitlement"
	catalogsourceblueprint "github.com/ankasoftco/provider-vra/internal/controller/catalogsourceblueprint/catalogsourceblueprint"
	catalogsourceentitlement "github.com/ankasoftco/provider-vra/internal/controller/catalogsourceentitlement/catalogsourceentitlement"
	accountaws "github.com/ankasoftco/provider-vra/internal/controller/cloudaccount/accountaws"
	accountazure "github.com/ankasoftco/provider-vra/internal/controller/cloudaccount/accountazure"
	accountgcp "github.com/ankasoftco/provider-vra/internal/controller/cloudaccount/accountgcp"
	accountnsxt "github.com/ankasoftco/provider-vra/internal/controller/cloudaccount/accountnsxt"
	accountvmc "github.com/ankasoftco/provider-vra/internal/controller/cloudaccount/accountvmc"
	accountvsphere "github.com/ankasoftco/provider-vra/internal/controller/cloudaccount/accountvsphere"
	contentsource "github.com/ankasoftco/provider-vra/internal/controller/contentsource/contentsource"
	deployment "github.com/ankasoftco/provider-vra/internal/controller/deployment/deployment"
	compute "github.com/ankasoftco/provider-vra/internal/controller/fabric/compute"
	datastorevsphere "github.com/ankasoftco/provider-vra/internal/controller/fabric/datastorevsphere"
	networkvsphere "github.com/ankasoftco/provider-vra/internal/controller/fabric/networkvsphere"
	profile "github.com/ankasoftco/provider-vra/internal/controller/flavorprofile/profile"
	profileimageprofile "github.com/ankasoftco/provider-vra/internal/controller/imageprofile/profile"
	integration "github.com/ankasoftco/provider-vra/internal/controller/integration/integration"
	loadbalancer "github.com/ankasoftco/provider-vra/internal/controller/loadbalancer/loadbalancer"
	machine "github.com/ankasoftco/provider-vra/internal/controller/machine/machine"
	network "github.com/ankasoftco/provider-vra/internal/controller/network/network"
	networkiprange "github.com/ankasoftco/provider-vra/internal/controller/network/networkiprange"
	networkprofile "github.com/ankasoftco/provider-vra/internal/controller/network/networkprofile"
	project "github.com/ankasoftco/provider-vra/internal/controller/project/project"
	providerconfig "github.com/ankasoftco/provider-vra/internal/controller/providerconfig"
	profilestorage "github.com/ankasoftco/provider-vra/internal/controller/storage/profile"
	profileaws "github.com/ankasoftco/provider-vra/internal/controller/storage/profileaws"
	profileazure "github.com/ankasoftco/provider-vra/internal/controller/storage/profileazure"
	profilevsphere "github.com/ankasoftco/provider-vra/internal/controller/storage/profilevsphere"
	zone "github.com/ankasoftco/provider-vra/internal/controller/zone/zone"
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
		accountaws.Setup,
		accountazure.Setup,
		accountgcp.Setup,
		accountnsxt.Setup,
		accountvmc.Setup,
		accountvsphere.Setup,
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
		network.Setup,
		networkiprange.Setup,
		networkprofile.Setup,
		project.Setup,
		providerconfig.Setup,
		profilestorage.Setup,
		profileaws.Setup,
		profileazure.Setup,
		profilevsphere.Setup,
		zone.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

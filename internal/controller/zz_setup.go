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
	deployment "github.com/ankasoftco/upjet-provider-vra/internal/controller/deployment/deployment"
	compute "github.com/ankasoftco/upjet-provider-vra/internal/controller/fabric/compute"
	project "github.com/ankasoftco/upjet-provider-vra/internal/controller/project/project"
	providerconfig "github.com/ankasoftco/upjet-provider-vra/internal/controller/providerconfig"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		blockdevice.Setup,
		blockdevicesnapshot.Setup,
		blueprint.Setup,
		version.Setup,
		deployment.Setup,
		compute.Setup,
		project.Setup,
		providerconfig.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

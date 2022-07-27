/*
Copyright 2022 The ANKA SOFTWARE Authors.
*/

package deployment

import (
	"context"
	"fmt"

	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	"github.com/crossplane/crossplane-runtime/pkg/connection"
	"github.com/crossplane/crossplane-runtime/pkg/controller"
	"github.com/crossplane/crossplane-runtime/pkg/event"
	"github.com/crossplane/crossplane-runtime/pkg/meta"
	"github.com/crossplane/crossplane-runtime/pkg/ratelimiter"
	"github.com/crossplane/crossplane-runtime/pkg/reconciler/managed"
	"github.com/crossplane/crossplane-runtime/pkg/resource"

	sdkCatalogItems "github.com/vmware/vra-sdk-go/pkg/client/catalog_items"
	sdkDeployment "github.com/vmware/vra-sdk-go/pkg/client/deployments"

	"github.com/pkg/errors"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/crossplane/provider-vra/apis/deployment/v1alpha1"
	apisv1alpha1 "github.com/crossplane/provider-vra/apis/v1alpha1"
	"github.com/crossplane/provider-vra/internal/clients"
	"github.com/crossplane/provider-vra/internal/clients/catalogitem"
	"github.com/crossplane/provider-vra/internal/clients/deployment"
	"github.com/crossplane/provider-vra/internal/controller/features"
)

const (
	errNotDeployment = "managed resource is not a Deployment custom resource"
	errTrackPCUsage  = "cannot track ProviderConfig usage"

	errCreateFailed = "cannot create deployment with vRA API"
	// errGetFailed    = "cannot get deployment from vRA API"
	errDeleteFailed = "cannot delete deployment from vRA API"
)

// Setup adds a controller that reconciles Deployment managed resources.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	name := managed.ControllerName(v1alpha1.DeploymentGroupKind)

	cps := []managed.ConnectionPublisher{managed.NewAPISecretPublisher(mgr.GetClient(), mgr.GetScheme())}
	if o.Features.Enabled(features.EnableAlphaExternalSecretStores) {
		cps = append(cps, connection.NewDetailsManager(mgr.GetClient(), apisv1alpha1.StoreConfigGroupVersionKind))
	}

	r := managed.NewReconciler(mgr,
		resource.ManagedKind(v1alpha1.DeploymentGroupVersionKind),
		managed.WithExternalConnecter(&connector{
			kube:                    mgr.GetClient(),
			usage:                   resource.NewProviderConfigUsageTracker(mgr.GetClient(), &apisv1alpha1.ProviderConfigUsage{}),
			newServiceFnDeployment:  deployment.NewDeploymentClient,
			newServiceFnCatalogItem: catalogitem.NewCatalogItemClient}),
		managed.WithLogger(o.Logger.WithValues("controller", name)),
		managed.WithRecorder(event.NewAPIRecorder(mgr.GetEventRecorderFor(name))),
		managed.WithConnectionPublishers(cps...))

	return ctrl.NewControllerManagedBy(mgr).
		Named(name).
		WithOptions(o.ForControllerRuntime()).
		For(&v1alpha1.Deployment{}).
		Complete(ratelimiter.NewReconciler(name, r, o.GlobalRateLimiter))
}

// A connector is expected to produce an ExternalClient when its Connect method
// is called.
type connector struct {
	kube                    client.Client
	usage                   resource.Tracker
	newServiceFnDeployment  func(cfg clients.Config) sdkDeployment.ClientService
	newServiceFnCatalogItem func(cfg clients.Config) sdkCatalogItems.ClientService
}

// Connect typically produces an ExternalClient by:
// 1. Tracking that the managed resource is using a ProviderConfig.
// 2. Getting the managed resource's ProviderConfig.
// 3. Getting the credentials specified by the ProviderConfig.
// 4. Using the credentials to form a client.
func (c *connector) Connect(ctx context.Context, mg resource.Managed) (managed.ExternalClient, error) {
	cr, ok := mg.(*v1alpha1.Deployment)
	if !ok {
		return nil, errors.New(errNotDeployment)
	}

	if err := c.usage.Track(ctx, mg); err != nil {
		return nil, errors.Wrap(err, errTrackPCUsage)
	}

	cfg, err := clients.GetConfig(ctx, c.kube, cr)
	if err != nil {
		return nil, err
	}

	return &external{kube: c.kube, serviceDeployment: c.newServiceFnDeployment(*cfg), serviceCatalogItem: c.newServiceFnCatalogItem(*cfg)}, nil
}

// An ExternalClient observes, then either creates, updates, or deletes an
// external resource to ensure it reflects the managed resource's desired state.
type external struct {
	kube client.Client
	// A 'client' used to connect to the external resource API.
	serviceDeployment  sdkDeployment.ClientService
	serviceCatalogItem sdkCatalogItems.ClientService
}

func (c *external) Observe(ctx context.Context, mg resource.Managed) (managed.ExternalObservation, error) {
	cr, ok := mg.(*v1alpha1.Deployment)
	if !ok {
		return managed.ExternalObservation{}, errors.New(errNotDeployment)
	}

	// These fmt statements should be removed in the real implementation.
	fmt.Printf("Observing: \n%+v", cr)

	externalName := meta.GetExternalName(cr)
	if externalName == "" {
		return managed.ExternalObservation{ResourceExists: false}, nil
	}

	deploymentID := externalName

	params := deployment.GenerateGetDeploymentOptions(deploymentID)
	deploymentObj, _ := c.serviceDeployment.GetDeploymentByIDV3UsingGET(params)

	if deploymentObj == nil {
		return managed.ExternalObservation{ResourceExists: false}, nil
	}

	// TODO: Check the deployment process here...
	if deploymentObj.Payload.Status == "CREATE_SUCCESSFUL" {
		cr.Status.SetConditions(xpv1.Available())
	} else if deploymentObj.Payload.Status == "CREATE_INPROGRESS" {
		cr.Status.SetConditions(xpv1.Creating())
	} else if deploymentObj.Payload.Status == "CREATE_FAILED" {
		return managed.ExternalObservation{}, nil
	} else if deploymentObj.Payload.Status == "DELETE_SUCCESSFUL" {
		return managed.ExternalObservation{}, nil
	} else if deploymentObj.Payload.Status == "DELETE_INPROGRESS" {
		cr.Status.SetConditions(xpv1.Deleting())
	}
	cr.Status.AtProvider = deployment.GenerateDeploymentObservation(deploymentObj)

	// cr.Status.SetConditions(xpv1.Available())
	// cr.Status.SetConditions(xpv1.Creating())

	return managed.ExternalObservation{
		// Return false when the external resource does not exist. This lets
		// the managed resource reconciler know that it needs to call Create to
		// (re)create the resource, or that it has successfully been deleted.
		ResourceExists: true,

		// Return false when the external resource exists, but it not up to date
		// with the desired managed resource state. This lets the managed
		// resource reconciler know that it needs to call Update.
		ResourceUpToDate: true,
	}, nil
}

func (c *external) Create(ctx context.Context, mg resource.Managed) (managed.ExternalCreation, error) {
	cr, ok := mg.(*v1alpha1.Deployment)
	if !ok {
		return managed.ExternalCreation{}, errors.New(errNotDeployment)
	}

	fmt.Printf("Creating: \n%+v", cr)

	// cr.Status.SetConditions(xpv1.Creating())

	deploymentParams := catalogitem.GenerateRequestCatalogItemOptions(&cr.Spec.ForProvider)

	response, err := c.serviceCatalogItem.RequestCatalogItemInstancesUsingPOST1(deploymentParams)
	if err != nil {
		return managed.ExternalCreation{}, errors.Wrap(err, errCreateFailed)
	}

	fmt.Println("Dep Id: " + response.Payload[0].DeploymentID)
	meta.SetExternalName(cr, response.Payload[0].DeploymentID)

	// cr.Status.SetConditions(xpv1.Available())
	cr.Status.SetConditions(xpv1.Creating())

	return managed.ExternalCreation{
		// Optionally return any details that may be required to connect to the
		// external resource. These will be stored as the connection secret.
		ConnectionDetails: managed.ConnectionDetails{},
	}, nil
}

func (c *external) Update(ctx context.Context, mg resource.Managed) (managed.ExternalUpdate, error) {
	cr, ok := mg.(*v1alpha1.Deployment)
	if !ok {
		return managed.ExternalUpdate{}, errors.New(errNotDeployment)
	}

	fmt.Printf("Updating: \n%+v", cr)

	cr.Status.SetConditions(xpv1.Available())

	return managed.ExternalUpdate{
		// Optionally return any details that may be required to connect to the
		// external resource. These will be stored as the connection secret.
		ConnectionDetails: managed.ConnectionDetails{},
	}, nil
}

func (c *external) Delete(ctx context.Context, mg resource.Managed) error {
	cr, ok := mg.(*v1alpha1.Deployment)
	if !ok {
		return errors.New(errNotDeployment)
	}

	fmt.Printf("Deleting: \n%+v", cr)

	externalName := meta.GetExternalName(cr)

	params := deployment.GenerateDeleteDeploymentOptions(externalName)
	if _, err := c.serviceDeployment.DeleteDeploymentUsingDELETE2(params); err != nil {
		return errors.Wrap(err, errDeleteFailed)
	}
	// todo: check the decommissioning process here...

	cr.Status.SetConditions(xpv1.Deleting())

	return nil
}

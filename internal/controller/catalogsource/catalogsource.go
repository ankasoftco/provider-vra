/*
Copyright 2022 The Crossplane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package catalogsource

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/crossplane/crossplane-runtime/pkg/connection"
	"github.com/crossplane/crossplane-runtime/pkg/controller"
	"github.com/crossplane/crossplane-runtime/pkg/event"
	"github.com/crossplane/crossplane-runtime/pkg/meta"
	"github.com/crossplane/crossplane-runtime/pkg/ratelimiter"
	"github.com/crossplane/crossplane-runtime/pkg/reconciler/managed"
	"github.com/crossplane/crossplane-runtime/pkg/resource"

	apisv1alpha1 "github.com/crossplane/provider-vraprovider/apis/v1alpha1"
	"github.com/crossplane/provider-vraprovider/apis/vra/v1alpha1"
	"github.com/crossplane/provider-vraprovider/internal/features"
	sdkCatalogSources "github.com/vmware/vra-sdk-go/pkg/client/catalog_sources"

	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	"github.com/crossplane/provider-vraprovider/internal/clients"
	catalog_source "github.com/crossplane/provider-vraprovider/internal/clients/catalog_source"
	catalog_sources "github.com/crossplane/provider-vraprovider/internal/clients/catalog_source"
)

const (
	errNotCatalogSource = "managed resource is not a CatalogSource custom resource"
	errTrackPCUsage     = "cannot track ProviderConfig usage"
	errGetPC            = "cannot get ProviderConfig"
	errGetCreds         = "cannot get credentials"
	errCreateFailed     = "cannot create Catalog Source with vRA API"
	errDeleteFailed     = "cannot delete Catalog Source from vRA API"
	errUpdateFailed     = "cannot update Catalog Source from vRA API"

	errNewClient = "cannot create new Service"
)

// A NoOpService does nothing.
type NoOpService struct{}

var (
	newNoOpService = func(_ []byte) (interface{}, error) { return &NoOpService{}, nil }
)

// Setup adds a controller that reconciles CatalogSources managed resources.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	name := managed.ControllerName(v1alpha1.CatalogSourceGroupKind)

	cps := []managed.ConnectionPublisher{managed.NewAPISecretPublisher(mgr.GetClient(), mgr.GetScheme())}
	if o.Features.Enabled(features.EnableAlphaExternalSecretStores) {
		cps = append(cps, connection.NewDetailsManager(mgr.GetClient(), apisv1alpha1.StoreConfigGroupVersionKind))
	}

	r := managed.NewReconciler(mgr,
		resource.ManagedKind(v1alpha1.CatalogSourceGroupVersionKind),
		managed.WithExternalConnecter(&connector{
			kube:         mgr.GetClient(),
			usage:        resource.NewProviderConfigUsageTracker(mgr.GetClient(), &apisv1alpha1.ProviderConfigUsage{}),
			newServiceFn: catalog_source.NewCatalogSourceClient}),
		managed.WithLogger(o.Logger.WithValues("controller", name)),
		managed.WithPollInterval(o.PollInterval),
		managed.WithRecorder(event.NewAPIRecorder(mgr.GetEventRecorderFor(name))),
		managed.WithConnectionPublishers(cps...))

	return ctrl.NewControllerManagedBy(mgr).
		Named(name).
		WithOptions(o.ForControllerRuntime()).
		WithEventFilter(resource.DesiredStateChanged()).
		For(&v1alpha1.CatalogSource{}).
		Complete(ratelimiter.NewReconciler(name, r, o.GlobalRateLimiter))
}

// A connector is expected to produce an ExternalClient when its Connect method
// is called.
type connector struct {
	kube         client.Client
	usage        resource.Tracker
	newServiceFn func(cfg clients.Config) sdkCatalogSources.ClientService
}

// Connect typically produces an ExternalClient by:
// 1. Tracking that the managed resource is using a ProviderConfig.
// 2. Getting the managed resource's ProviderConfig.
// 3. Getting the credentials specified by the ProviderConfig.
// 4. Using the credentials to form a client.
func (c *connector) Connect(ctx context.Context, mg resource.Managed) (managed.ExternalClient, error) {
	cr, ok := mg.(*v1alpha1.CatalogSource)
	if !ok {
		return nil, errors.New(errNotCatalogSource)
	}

	if err := c.usage.Track(ctx, mg); err != nil {
		return nil, errors.Wrap(err, errTrackPCUsage)
	}

	cfg, err := clients.GetConfig(ctx, c.kube, cr)
	if err != nil {
		return nil, err
	}
	return &external{kube: c.kube, service: c.newServiceFn(*cfg)}, nil
}

// An ExternalClient observes, then either creates, updates, or deletes an
// external resource to ensure it reflects the managed resource's desired state.
type external struct {
	kube client.Client
	// A 'client' used to connect to the external resource API. In practice this
	// would be something like an AWS SDK client.
	service sdkCatalogSources.ClientService
}

func (c *external) Observe(ctx context.Context, mg resource.Managed) (managed.ExternalObservation, error) {
	cr, ok := mg.(*v1alpha1.CatalogSource)
	if !ok {
		return managed.ExternalObservation{}, errors.New(errNotCatalogSource)
	}

	// These fmt statements should be removed in the real implementation.
	fmt.Printf("Observing: %+v", cr)

	externalName := meta.GetExternalName(cr)
	if externalName == "" {
		return managed.ExternalObservation{ResourceExists: false}, nil
	}

	catalog_source_ID := externalName

	params := catalog_source.GenerateGetCatalogSourceOptions(catalog_source_ID)

	// current value of the resource.
	b, _ := c.service.GetUsingGET2(params)

	// desired value of the resource. (it reads the current yaml to determine desired state)
	desired := cr.Spec.ForProvider.DeepCopy()

	if b == nil {
		return managed.ExternalObservation{ResourceExists: false}, nil
	}

	ResourceUpToDate := catalog_source.IsResourceUpToDate(desired, b.Payload)

	// TODO: Check the deployment process here...
	cr.Status.AtProvider = catalog_source.GenerateCatalogSourceObservation(b)
	cr.Status.SetConditions(xpv1.Available())

	return managed.ExternalObservation{
		// Return false when the external resource does not exist. This lets
		// the managed resource reconciler know that it needs to call Create to
		// (re)create the resource, or that it has successfully been deleted.
		ResourceExists: true,

		// Return false when the external resource exists, but it not up to date
		// with the desired managed resource state. This lets the managed
		// resource reconciler know that it needs to call Update.
		ResourceUpToDate: ResourceUpToDate,

		// Return any details that may be required to connect to the external
		// resource. These will be stored as the connection secret.
		ConnectionDetails: managed.ConnectionDetails{},
	}, nil
}

func (c *external) Create(ctx context.Context, mg resource.Managed) (managed.ExternalCreation, error) {
	cr, ok := mg.(*v1alpha1.CatalogSource)
	if !ok {
		return managed.ExternalCreation{}, errors.New(errNotCatalogSource)
	}

	fmt.Printf("Creating: %+v", cr)

	cr.Status.SetConditions(xpv1.Creating())

	catalogSourceParam := catalog_sources.GenerateCreateCatalogSourceOptions(&cr.Spec.ForProvider)

	responseOk, responseCreated, err := c.service.PostUsingPOST2(catalogSourceParam)
	if err != nil {
		return managed.ExternalCreation{}, errors.Wrap(err, errCreateFailed)
	}

	if responseOk != nil {
		fmt.Print("Catalog Source Id: ")
		fmt.Println(responseOk.Payload.ID)
		meta.SetExternalName(cr, fmt.Sprint(*&responseOk.Payload.ID))
	}
	if responseCreated != nil {
		fmt.Print("Catalog Source Id: ")
		fmt.Println(responseCreated.Payload.ID)
		meta.SetExternalName(cr, fmt.Sprint(*&responseCreated.Payload.ID))
	}

	cr.Status.SetConditions(xpv1.Available())
	return managed.ExternalCreation{
		// Optionally return any details that may be required to connect to the
		// external resource. These will be stored as the connection secret.
		ConnectionDetails: managed.ConnectionDetails{},
	}, nil
}

func (c *external) Update(ctx context.Context, mg resource.Managed) (managed.ExternalUpdate, error) {
	cr, ok := mg.(*v1alpha1.CatalogSource)
	if !ok {
		return managed.ExternalUpdate{}, errors.New(errNotCatalogSource)
	}

	fmt.Printf("Updating: %+v", cr)

	externalName := meta.GetExternalName(cr)

	catalogSourceParam := catalog_sources.GenerateUpdateCatalogSourceOptions(externalName, &cr.Spec.ForProvider)
	responseOk, responseCreated, err := c.service.PostUsingPOST2(catalogSourceParam)
	if err != nil {
		return managed.ExternalUpdate{}, errors.Wrap(err, errCreateFailed)
	}

	if responseOk != nil {
		fmt.Print("Catalog Source Id: ")
		fmt.Println(responseOk.Payload.ID)
		meta.SetExternalName(cr, fmt.Sprint(*&responseOk.Payload.ID))
	}
	if responseCreated != nil {
		fmt.Print("Catalog Source Id: ")
		fmt.Println(responseCreated.Payload.ID)
		meta.SetExternalName(cr, fmt.Sprint(*&responseCreated.Payload.ID))
	}
	cr.Status.SetConditions(xpv1.Available())

	return managed.ExternalUpdate{
		// Optionally return any details that may be required to connect to the
		// external resource. These will be stored as the connection secret.
		ConnectionDetails: managed.ConnectionDetails{},
	}, nil
}

func (c *external) Delete(ctx context.Context, mg resource.Managed) error {
	cr, ok := mg.(*v1alpha1.CatalogSource)
	if !ok {
		return errors.New(errNotCatalogSource)
	}

	fmt.Printf("Deleting: %+v", cr)

	externalName := meta.GetExternalName(cr)

	blueprintParam := catalog_source.GenerateDeleteCatalogSourceOptions(externalName)

	if _, err := c.service.DeleteUsingDELETE4(blueprintParam); err != nil {
		return errors.Wrap(err, errDeleteFailed)
	}

	cr.Status.SetConditions(xpv1.Available())

	return nil
}

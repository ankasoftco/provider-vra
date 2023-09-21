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

package deploymentactions

import (
	"context"
	"fmt"

	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
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
	sdkDeploymentActions "github.com/vmware/vra-sdk-go/pkg/client/deployment_actions"

	apisv1alpha1 "github.com/crossplane/provider-vraprovider/apis/v1alpha1"
	"github.com/crossplane/provider-vraprovider/apis/vra/v1alpha1"
	"github.com/crossplane/provider-vraprovider/internal/clients"
	deployment_actions "github.com/crossplane/provider-vraprovider/internal/clients/deployment_actions"
	"github.com/crossplane/provider-vraprovider/internal/features"
)

const (
	errNotDeploymentActions = "managed resource is not a DeploymentActions custom resource"
	errTrackPCUsage         = "cannot track ProviderConfig usage"
	errGetPC                = "cannot get ProviderConfig"
	errGetCreds             = "cannot get credentials"

	errCreateFailed = "cannot create Deployment Action with vRA API"
	errDeleteFailed = "cannot delete Deployment Action from vRA API"
	errUpdateFailed = "cannot update Deployment Action from vRA API"

	errNewClient = "cannot create new Service"
)

// A NoOpService does nothing.
type NoOpService struct{}

var (
	newNoOpService = func(_ []byte) (interface{}, error) { return &NoOpService{}, nil }
)

// Setup adds a controller that reconciles DeploymentActions managed resources.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	name := managed.ControllerName(v1alpha1.DeploymentActionsGroupKind)

	cps := []managed.ConnectionPublisher{managed.NewAPISecretPublisher(mgr.GetClient(), mgr.GetScheme())}
	if o.Features.Enabled(features.EnableAlphaExternalSecretStores) {
		cps = append(cps, connection.NewDetailsManager(mgr.GetClient(), apisv1alpha1.StoreConfigGroupVersionKind))
	}

	r := managed.NewReconciler(mgr,
		resource.ManagedKind(v1alpha1.DeploymentActionsGroupVersionKind),
		managed.WithExternalConnecter(&connector{
			kube:         mgr.GetClient(),
			usage:        resource.NewProviderConfigUsageTracker(mgr.GetClient(), &apisv1alpha1.ProviderConfigUsage{}),
			newServiceFn: deployment_actions.DeploymentActionsClient}),
		managed.WithLogger(o.Logger.WithValues("controller", name)),
		managed.WithPollInterval(o.PollInterval),
		managed.WithRecorder(event.NewAPIRecorder(mgr.GetEventRecorderFor(name))),
		managed.WithConnectionPublishers(cps...))

	return ctrl.NewControllerManagedBy(mgr).
		Named(name).
		WithOptions(o.ForControllerRuntime()).
		WithEventFilter(resource.DesiredStateChanged()).
		For(&v1alpha1.DeploymentActions{}).
		Complete(ratelimiter.NewReconciler(name, r, o.GlobalRateLimiter))
}

// A connector is expected to produce an ExternalClient when its Connect method
// is called.
type connector struct {
	kube         client.Client
	usage        resource.Tracker
	newServiceFn func(cfg clients.Config) sdkDeploymentActions.ClientService
}

// Connect typically produces an ExternalClient by:
// 1. Tracking that the managed resource is using a ProviderConfig.
// 2. Getting the managed resource's ProviderConfig.
// 3. Getting the credentials specified by the ProviderConfig.
// 4. Using the credentials to form a client.
func (c *connector) Connect(ctx context.Context, mg resource.Managed) (managed.ExternalClient, error) {
	cr, ok := mg.(*v1alpha1.DeploymentActions)
	if !ok {
		return nil, errors.New(errNotDeploymentActions)
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
	service sdkDeploymentActions.ClientService
}

func (c *external) Observe(ctx context.Context, mg resource.Managed) (managed.ExternalObservation, error) {
	cr, ok := mg.(*v1alpha1.DeploymentActions)
	if !ok {
		return managed.ExternalObservation{}, errors.New(errNotDeploymentActions)
	}
	// These fmt statements should be removed in the real implementation.
	fmt.Printf("Observing: %+v", cr)

	externalName := meta.GetExternalName(cr)
	if externalName == "" {
		return managed.ExternalObservation{ResourceExists: false}, nil // trigger Create
	}
	// TODO: bug.. deployment id & name

	deploymentActionID := externalName

	params := deployment_actions.GenerateGetDeploymentActionsOptions(deploymentActionID, &cr.Spec.ForProvider)
	deployment_action_Obj, _ := c.service.GetDeploymentActionUsingGET2(params)

	if deployment_action_Obj == nil {
		return managed.ExternalObservation{ResourceExists: false}, nil // trigger Create
	}

	for _, element := range cr.Status.Conditions {
		if element.Reason == "Deleting" {
			return managed.ExternalObservation{ResourceExists: false}, nil // Deleting
		}
	}

	cr.Status.AtProvider = deployment_actions.GenerateDeploymentActionsObservation(deployment_action_Obj)
	cr.Status.SetConditions(xpv1.Available())

	return managed.ExternalObservation{
		// Return false when the external resource does not exist. This lets
		// the managed resource reconciler know that it needs to call Create to
		// (re)create the resource, or that it has successfully been deleted.
		ResourceExists: true,

		// Return false when the external resource exists, but it not up to date
		// with the desired managed resource state. This lets the managed
		// resource reconciler know that it needs to call Update.
		ResourceUpToDate: true,

		// Return any details that may be required to connect to the external
		// resource. These will be stored as the connection secret.
		ConnectionDetails: managed.ConnectionDetails{},
	}, nil
}

func (c *external) Create(ctx context.Context, mg resource.Managed) (managed.ExternalCreation, error) {
	cr, ok := mg.(*v1alpha1.DeploymentActions)
	if !ok {
		return managed.ExternalCreation{}, errors.New(errNotDeploymentActions)
	}

	fmt.Printf("Creating: %+v", cr)
	cr.Status.SetConditions(xpv1.Creating())

	deploymentActionsParams := deployment_actions.GenerateCreateDeploymentActionsOptions(&cr.Spec.ForProvider)

	response, err := c.service.SubmitDeploymentActionRequestUsingPOST2(deploymentActionsParams)
	if err != nil {
		return managed.ExternalCreation{}, errors.Wrap(err, errCreateFailed)
	}

	fmt.Println("Deployment Action Id: " + *&response.Payload.ActionID)
	meta.SetExternalName(cr, fmt.Sprint(*&response.Payload.ActionID))

	cr.Status.SetConditions(xpv1.Available())
	return managed.ExternalCreation{
		// Optionally return any details that may be required to connect to the
		// external resource. These will be stored as the connection secret.
		ConnectionDetails: managed.ConnectionDetails{},
	}, nil
}

func (c *external) Update(ctx context.Context, mg resource.Managed) (managed.ExternalUpdate, error) {
	cr, ok := mg.(*v1alpha1.DeploymentActions)
	if !ok {
		return managed.ExternalUpdate{}, errors.New(errNotDeploymentActions)
	}

	fmt.Printf("Updating: %+v", cr)

	return managed.ExternalUpdate{
		// Optionally return any details that may be required to connect to the
		// external resource. These will be stored as the connection secret.
		ConnectionDetails: managed.ConnectionDetails{},
	}, nil
}

func (c *external) Delete(ctx context.Context, mg resource.Managed) error {
	cr, ok := mg.(*v1alpha1.DeploymentActions)
	if !ok {
		return errors.New(errNotDeploymentActions)
	}

	fmt.Printf("Deleting: %+v", cr)

	cr.Status.SetConditions(xpv1.Deleting())

	return nil
}

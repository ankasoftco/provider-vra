/*
Copyright 2023 Upbound Inc. - ANKASOFT
*/

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type HealthCheckConfigurationObservation struct {
	HealthyThreshold *float64 `json:"healthyThreshold,omitempty" tf:"healthy_threshold,omitempty"`

	IntervalSeconds *float64 `json:"intervalSeconds,omitempty" tf:"interval_seconds,omitempty"`

	Port *string `json:"port,omitempty" tf:"port,omitempty"`

	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`

	TimeoutSeconds *float64 `json:"timeoutSeconds,omitempty" tf:"timeout_seconds,omitempty"`

	URLPath *string `json:"urlPath,omitempty" tf:"url_path,omitempty"`

	UnhealthyThreshold *float64 `json:"unhealthyThreshold,omitempty" tf:"unhealthy_threshold,omitempty"`
}

type HealthCheckConfigurationParameters struct {

	// +kubebuilder:validation:Optional
	HealthyThreshold *float64 `json:"healthyThreshold,omitempty" tf:"healthy_threshold,omitempty"`

	// +kubebuilder:validation:Optional
	IntervalSeconds *float64 `json:"intervalSeconds,omitempty" tf:"interval_seconds,omitempty"`

	// +kubebuilder:validation:Required
	Port *string `json:"port" tf:"port,omitempty"`

	// +kubebuilder:validation:Required
	Protocol *string `json:"protocol" tf:"protocol,omitempty"`

	// +kubebuilder:validation:Optional
	TimeoutSeconds *float64 `json:"timeoutSeconds,omitempty" tf:"timeout_seconds,omitempty"`

	// +kubebuilder:validation:Optional
	URLPath *string `json:"urlPath,omitempty" tf:"url_path,omitempty"`

	// +kubebuilder:validation:Optional
	UnhealthyThreshold *float64 `json:"unhealthyThreshold,omitempty" tf:"unhealthy_threshold,omitempty"`
}

type LinksObservation struct {
	Href *string `json:"href,omitempty" tf:"href,omitempty"`

	Hrefs []*string `json:"hrefs,omitempty" tf:"hrefs,omitempty"`

	Rel *string `json:"rel,omitempty" tf:"rel,omitempty"`
}

type LinksParameters struct {
}

type LoadBalancerObservation struct {
	Address *string `json:"address,omitempty" tf:"address,omitempty"`

	CreatedAt *string `json:"createdAt,omitempty" tf:"created_at,omitempty"`

	CustomProperties map[string]*string `json:"customProperties,omitempty" tf:"custom_properties,omitempty"`

	DeploymentID *string `json:"deploymentId,omitempty" tf:"deployment_id,omitempty"`

	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	ExternalID *string `json:"externalId,omitempty" tf:"external_id,omitempty"`

	ExternalRegionID *string `json:"externalRegionId,omitempty" tf:"external_region_id,omitempty"`

	ExternalZoneID *string `json:"externalZoneId,omitempty" tf:"external_zone_id,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	InternetFacing *bool `json:"internetFacing,omitempty" tf:"internet_facing,omitempty"`

	Links []LinksObservation `json:"links,omitempty" tf:"links,omitempty"`

	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	Nics []NicsObservation `json:"nics,omitempty" tf:"nics,omitempty"`

	OrganizationID *string `json:"organizationId,omitempty" tf:"organization_id,omitempty"`

	Owner *string `json:"owner,omitempty" tf:"owner,omitempty"`

	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`

	Routes []RoutesObservation `json:"routes,omitempty" tf:"routes,omitempty"`

	Tags []TagsObservation `json:"tags,omitempty" tf:"tags,omitempty"`

	Targets []TargetsObservation `json:"targets,omitempty" tf:"targets,omitempty"`

	UpdatedAt *string `json:"updatedAt,omitempty" tf:"updated_at,omitempty"`
}

type LoadBalancerParameters struct {

	// +kubebuilder:validation:Optional
	CustomProperties map[string]*string `json:"customProperties,omitempty" tf:"custom_properties,omitempty"`

	// +kubebuilder:validation:Optional
	DeploymentID *string `json:"deploymentId,omitempty" tf:"deployment_id,omitempty"`

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Optional
	InternetFacing *bool `json:"internetFacing,omitempty" tf:"internet_facing,omitempty"`

	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	Nics []NicsParameters `json:"nics,omitempty" tf:"nics,omitempty"`

	// +crossplane:generate:reference:type=github.com/ankasoftco/provider-vra/apis/project/v1alpha1.Project
	// +kubebuilder:validation:Optional
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// Reference to a Project in project to populate projectId.
	// +kubebuilder:validation:Optional
	ProjectIDRef *v1.Reference `json:"projectIdRef,omitempty" tf:"-"`

	// Selector for a Project in project to populate projectId.
	// +kubebuilder:validation:Optional
	ProjectIDSelector *v1.Selector `json:"projectIdSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	Routes []RoutesParameters `json:"routes,omitempty" tf:"routes,omitempty"`

	// +kubebuilder:validation:Optional
	Tags []TagsParameters `json:"tags,omitempty" tf:"tags,omitempty"`

	// +kubebuilder:validation:Optional
	Targets []TargetsParameters `json:"targets,omitempty" tf:"targets,omitempty"`
}

type NicsObservation struct {
	Addresses []*string `json:"addresses,omitempty" tf:"addresses,omitempty"`

	CustomProperties map[string]*string `json:"customProperties,omitempty" tf:"custom_properties,omitempty"`

	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	DeviceIndex *float64 `json:"deviceIndex,omitempty" tf:"device_index,omitempty"`

	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	NetworkID *string `json:"networkId,omitempty" tf:"network_id,omitempty"`

	SecurityGroupIds []*string `json:"securityGroupIds,omitempty" tf:"security_group_ids,omitempty"`
}

type NicsParameters struct {

	// +kubebuilder:validation:Optional
	Addresses []*string `json:"addresses,omitempty" tf:"addresses,omitempty"`

	// +kubebuilder:validation:Optional
	CustomProperties map[string]*string `json:"customProperties,omitempty" tf:"custom_properties,omitempty"`

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Optional
	DeviceIndex *float64 `json:"deviceIndex,omitempty" tf:"device_index,omitempty"`

	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	NetworkID *string `json:"networkId" tf:"network_id,omitempty"`

	// +kubebuilder:validation:Optional
	SecurityGroupIds []*string `json:"securityGroupIds,omitempty" tf:"security_group_ids,omitempty"`
}

type RoutesObservation struct {
	HealthCheckConfiguration []HealthCheckConfigurationObservation `json:"healthCheckConfiguration,omitempty" tf:"health_check_configuration,omitempty"`

	MemberPort *string `json:"memberPort,omitempty" tf:"member_port,omitempty"`

	MemberProtocol *string `json:"memberProtocol,omitempty" tf:"member_protocol,omitempty"`

	Port *string `json:"port,omitempty" tf:"port,omitempty"`

	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`
}

type RoutesParameters struct {

	// +kubebuilder:validation:Optional
	HealthCheckConfiguration []HealthCheckConfigurationParameters `json:"healthCheckConfiguration,omitempty" tf:"health_check_configuration,omitempty"`

	// +kubebuilder:validation:Required
	MemberPort *string `json:"memberPort" tf:"member_port,omitempty"`

	// +kubebuilder:validation:Required
	MemberProtocol *string `json:"memberProtocol" tf:"member_protocol,omitempty"`

	// +kubebuilder:validation:Required
	Port *string `json:"port" tf:"port,omitempty"`

	// +kubebuilder:validation:Required
	Protocol *string `json:"protocol" tf:"protocol,omitempty"`
}

type TagsObservation struct {
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type TagsParameters struct {

	// +kubebuilder:validation:Required
	Key *string `json:"key" tf:"key,omitempty"`

	// +kubebuilder:validation:Required
	Value *string `json:"value" tf:"value,omitempty"`
}

type TargetsObservation struct {
	MachineID *string `json:"machineId,omitempty" tf:"machine_id,omitempty"`

	NetworkInterfaceID *string `json:"networkInterfaceId,omitempty" tf:"network_interface_id,omitempty"`
}

type TargetsParameters struct {

	// +kubebuilder:validation:Required
	MachineID *string `json:"machineId" tf:"machine_id,omitempty"`

	// +kubebuilder:validation:Optional
	NetworkInterfaceID *string `json:"networkInterfaceId,omitempty" tf:"network_interface_id,omitempty"`
}

// LoadBalancerSpec defines the desired state of LoadBalancer
type LoadBalancerSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     LoadBalancerParameters `json:"forProvider"`
}

// LoadBalancerStatus defines the observed state of LoadBalancer.
type LoadBalancerStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        LoadBalancerObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// LoadBalancer is the Schema for the LoadBalancers API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,vra}
type LoadBalancer struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.name)",message="name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.nics)",message="nics is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.routes)",message="routes is a required parameter"
	Spec   LoadBalancerSpec   `json:"spec"`
	Status LoadBalancerStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// LoadBalancerList contains a list of LoadBalancers
type LoadBalancerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LoadBalancer `json:"items"`
}

// Repository type metadata.
var (
	LoadBalancer_Kind             = "LoadBalancer"
	LoadBalancer_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: LoadBalancer_Kind}.String()
	LoadBalancer_KindAPIVersion   = LoadBalancer_Kind + "." + CRDGroupVersion.String()
	LoadBalancer_GroupVersionKind = CRDGroupVersion.WithKind(LoadBalancer_Kind)
)

func init() {
	SchemeBuilder.Register(&LoadBalancer{}, &LoadBalancerList{})
}

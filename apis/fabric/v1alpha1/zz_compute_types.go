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

type ComputeObservation struct {

	// Date when the entity was created. The date is in ISO 8601 and UTC.
	CreatedAt *string `json:"createdAt,omitempty" tf:"created_at,omitempty"`

	// A list of key value pair of custom properties for the fabric compute resource.
	CustomProperties map[string]*string `json:"customProperties,omitempty" tf:"custom_properties,omitempty"`

	// A human-friendly description.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The id of the external entity on the provider side.
	ExternalID *string `json:"externalId,omitempty" tf:"external_id,omitempty"`

	// The external region id of the fabric compute.
	ExternalRegionID *string `json:"externalRegionId,omitempty" tf:"external_region_id,omitempty"`

	// The external zone id of the fabric compute.
	ExternalZoneID *string `json:"externalZoneId,omitempty" tf:"external_zone_id,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Lifecycle status of the compute instance.
	LifecycleState *string `json:"lifecycleState,omitempty" tf:"lifecycle_state,omitempty"`

	Links []LinksObservation `json:"links,omitempty" tf:"links,omitempty"`

	// A human-friendly name used as an identifier for the fabric compute resource instance.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The id of the organization this entity belongs to.
	OrgID *string `json:"orgId,omitempty" tf:"org_id,omitempty"`

	// Email of the user that owns the entity.
	Owner *string `json:"owner,omitempty" tf:"owner,omitempty"`

	// Power state of fabric compute instance.
	PowerState *string `json:"powerState,omitempty" tf:"power_state,omitempty"`

	Tags []TagsObservation `json:"tags,omitempty" tf:"tags,omitempty"`

	// Type of the fabric compute instance.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// Date when the entity was last updated. The date is ISO 8601 and UTC.
	UpdatedAt *string `json:"updatedAt,omitempty" tf:"updated_at,omitempty"`
}

type ComputeParameters struct {

	// +kubebuilder:validation:Optional
	Tags []TagsParameters `json:"tags,omitempty" tf:"tags,omitempty"`
}

type LinksObservation struct {
	Href *string `json:"href,omitempty" tf:"href,omitempty"`

	Hrefs []*string `json:"hrefs,omitempty" tf:"hrefs,omitempty"`

	Rel *string `json:"rel,omitempty" tf:"rel,omitempty"`
}

type LinksParameters struct {
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

// ComputeSpec defines the desired state of Compute
type ComputeSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ComputeParameters `json:"forProvider"`
}

// ComputeStatus defines the observed state of Compute.
type ComputeStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ComputeObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Compute is the Schema for the Computes API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,vra}
type Compute struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ComputeSpec   `json:"spec"`
	Status            ComputeStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ComputeList contains a list of Computes
type ComputeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Compute `json:"items"`
}

// Repository type metadata.
var (
	Compute_Kind             = "Compute"
	Compute_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Compute_Kind}.String()
	Compute_KindAPIVersion   = Compute_Kind + "." + CRDGroupVersion.String()
	Compute_GroupVersionKind = CRDGroupVersion.WithKind(Compute_Kind)
)

func init() {
	SchemeBuilder.Register(&Compute{}, &ComputeList{})
}
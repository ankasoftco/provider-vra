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

type ProfileAwsLinksObservation struct {
	Href *string `json:"href,omitempty" tf:"href,omitempty"`

	Hrefs []*string `json:"hrefs,omitempty" tf:"hrefs,omitempty"`

	Rel *string `json:"rel,omitempty" tf:"rel,omitempty"`
}

type ProfileAwsLinksParameters struct {
}

type ProfileAwsObservation struct {
	CreatedAt *string `json:"createdAt,omitempty" tf:"created_at,omitempty"`

	DefaultItem *bool `json:"defaultItem,omitempty" tf:"default_item,omitempty"`

	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	DeviceType *string `json:"deviceType,omitempty" tf:"device_type,omitempty"`

	ExternalRegionID *string `json:"externalRegionId,omitempty" tf:"external_region_id,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	Iops *string `json:"iops,omitempty" tf:"iops,omitempty"`

	Links []ProfileAwsLinksObservation `json:"links,omitempty" tf:"links,omitempty"`

	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	OrganizationID *string `json:"organizationId,omitempty" tf:"organization_id,omitempty"`

	Owner *string `json:"owner,omitempty" tf:"owner,omitempty"`

	RegionID *string `json:"regionId,omitempty" tf:"region_id,omitempty"`

	SupportsEncryption *bool `json:"supportsEncryption,omitempty" tf:"supports_encryption,omitempty"`

	Tags []ProfileAwsTagsObservation `json:"tags,omitempty" tf:"tags,omitempty"`

	UpdatedAt *string `json:"updatedAt,omitempty" tf:"updated_at,omitempty"`

	VolumeType *string `json:"volumeType,omitempty" tf:"volume_type,omitempty"`
}

type ProfileAwsParameters struct {

	// +kubebuilder:validation:Optional
	DefaultItem *bool `json:"defaultItem,omitempty" tf:"default_item,omitempty"`

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Optional
	DeviceType *string `json:"deviceType,omitempty" tf:"device_type,omitempty"`

	// +kubebuilder:validation:Optional
	Iops *string `json:"iops,omitempty" tf:"iops,omitempty"`

	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	RegionID *string `json:"regionId,omitempty" tf:"region_id,omitempty"`

	// +kubebuilder:validation:Optional
	SupportsEncryption *bool `json:"supportsEncryption,omitempty" tf:"supports_encryption,omitempty"`

	// +kubebuilder:validation:Optional
	Tags []ProfileAwsTagsParameters `json:"tags,omitempty" tf:"tags,omitempty"`

	// +kubebuilder:validation:Optional
	VolumeType *string `json:"volumeType,omitempty" tf:"volume_type,omitempty"`
}

type ProfileAwsTagsObservation struct {
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type ProfileAwsTagsParameters struct {

	// +kubebuilder:validation:Required
	Key *string `json:"key" tf:"key,omitempty"`

	// +kubebuilder:validation:Required
	Value *string `json:"value" tf:"value,omitempty"`
}

// ProfileAwsSpec defines the desired state of ProfileAws
type ProfileAwsSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ProfileAwsParameters `json:"forProvider"`
}

// ProfileAwsStatus defines the observed state of ProfileAws.
type ProfileAwsStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ProfileAwsObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ProfileAws is the Schema for the ProfileAwss API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,vra}
type ProfileAws struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.defaultItem)",message="defaultItem is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.name)",message="name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.regionId)",message="regionId is a required parameter"
	Spec   ProfileAwsSpec   `json:"spec"`
	Status ProfileAwsStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ProfileAwsList contains a list of ProfileAwss
type ProfileAwsList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ProfileAws `json:"items"`
}

// Repository type metadata.
var (
	ProfileAws_Kind             = "ProfileAws"
	ProfileAws_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ProfileAws_Kind}.String()
	ProfileAws_KindAPIVersion   = ProfileAws_Kind + "." + CRDGroupVersion.String()
	ProfileAws_GroupVersionKind = CRDGroupVersion.WithKind(ProfileAws_Kind)
)

func init() {
	SchemeBuilder.Register(&ProfileAws{}, &ProfileAwsList{})
}

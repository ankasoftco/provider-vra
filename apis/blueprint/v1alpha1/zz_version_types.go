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

type VersionObservation struct {
	BlueprintDescription *string `json:"blueprintDescription,omitempty" tf:"blueprint_description,omitempty"`

	BlueprintID *string `json:"blueprintId,omitempty" tf:"blueprint_id,omitempty"`

	ChangeLog *string `json:"changeLog,omitempty" tf:"change_log,omitempty"`

	Content *string `json:"content,omitempty" tf:"content,omitempty"`

	CreatedAt *string `json:"createdAt,omitempty" tf:"created_at,omitempty"`

	CreatedBy *string `json:"createdBy,omitempty" tf:"created_by,omitempty"`

	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	OrgID *string `json:"orgId,omitempty" tf:"org_id,omitempty"`

	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`

	ProjectName *string `json:"projectName,omitempty" tf:"project_name,omitempty"`

	Release *bool `json:"release,omitempty" tf:"release,omitempty"`

	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	UpdatedAt *string `json:"updatedAt,omitempty" tf:"updated_at,omitempty"`

	UpdatedBy *string `json:"updatedBy,omitempty" tf:"updated_by,omitempty"`

	Valid *string `json:"valid,omitempty" tf:"valid,omitempty"`

	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

type VersionParameters struct {

	// +kubebuilder:validation:Optional
	BlueprintID *string `json:"blueprintId,omitempty" tf:"blueprint_id,omitempty"`

	// +kubebuilder:validation:Optional
	ChangeLog *string `json:"changeLog,omitempty" tf:"change_log,omitempty"`

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Optional
	Release *bool `json:"release,omitempty" tf:"release,omitempty"`

	// +kubebuilder:validation:Optional
	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

// VersionSpec defines the desired state of Version
type VersionSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     VersionParameters `json:"forProvider"`
}

// VersionStatus defines the observed state of Version.
type VersionStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        VersionObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Version is the Schema for the Versions API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,vra}
type Version struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.blueprintId)",message="blueprintId is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.version)",message="version is a required parameter"
	Spec   VersionSpec   `json:"spec"`
	Status VersionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VersionList contains a list of Versions
type VersionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Version `json:"items"`
}

// Repository type metadata.
var (
	Version_Kind             = "Version"
	Version_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Version_Kind}.String()
	Version_KindAPIVersion   = Version_Kind + "." + CRDGroupVersion.String()
	Version_GroupVersionKind = CRDGroupVersion.WithKind(Version_Kind)
)

func init() {
	SchemeBuilder.Register(&Version{}, &VersionList{})
}

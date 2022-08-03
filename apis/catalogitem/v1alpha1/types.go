/*
Copyright 2022 The ANKA SOFTWARE Authors.
*/

package v1alpha1

import (
	"reflect"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

// CatalogItemParameters are the configurable fields of a CatalogItem.
type CatalogItemParameters struct {
}

// CatalogItemObservation are the observable fields of a CatalogItem.
type CatalogItemObservation struct {
}

// A CatalogItemSpec defines the desired state of a CatalogItem.
type CatalogItemSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       CatalogItemParameters `json:"forProvider"`
}

// A CatalogItemStatus represents the observed state of a CatalogItemStatus.
type CatalogItemStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          CatalogItemObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// A CatalogItem is a vRA Catalog Item API type.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,vra}
type CatalogItem struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   CatalogItemSpec   `json:"spec"`
	Status CatalogItemStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CatalogItemList contains a list of Catalog Item
type CatalogItemList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CatalogItem `json:"items"`
}

// CatalogItem type metadata.
var (
	CatalogItemKind             = reflect.TypeOf(CatalogItem{}).Name()
	CatalogItemGroupKind        = schema.GroupKind{Group: Group, Kind: CatalogItemKind}.String()
	CatalogItemKindAPIVersion   = CatalogItemKind + "." + SchemeGroupVersion.String()
	CatalogItemGroupVersionKind = SchemeGroupVersion.WithKind(CatalogItemKind)
)

func init() {
	SchemeBuilder.Register(&CatalogItem{}, &CatalogItemList{})
}

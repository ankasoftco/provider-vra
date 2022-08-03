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

// LoginParameters are the configurable fields of a Login.
type LoginParameters struct {
	// Refresh token obtained from the UI
	// +immutable
	RefreshToken *string `json:"refreshToken"`
}

// LoginObservation are the observable fields of a Login.
type LoginObservation struct {
}

// A LoginSpec defines the desired state of a Login.
type LoginSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       LoginParameters `json:"forProvider"`
}

// A LoginStatus represents the observed state of a Login.
type LoginStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          LoginObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// A Login is an example API type.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,vra}
type Login struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   LoginSpec   `json:"spec"`
	Status LoginStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// LoginList contains a list of Login
type LoginList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Login `json:"items"`
}

// Login type metadata.
var (
	LoginKind             = reflect.TypeOf(Login{}).Name()
	LoginGroupKind        = schema.GroupKind{Group: Group, Kind: LoginKind}.String()
	LoginKindAPIVersion   = LoginKind + "." + SchemeGroupVersion.String()
	LoginGroupVersionKind = SchemeGroupVersion.WithKind(LoginKind)
)

func init() {
	SchemeBuilder.Register(&Login{}, &LoginList{})
}

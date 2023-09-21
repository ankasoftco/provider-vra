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

package v1alpha1

import (
	"reflect"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

// ResourceActionsParameters are the configurable fields of a ResourceActions.
type ResourceActionsParameters struct {
	// Resource resource id
	ResourceID string `json:"resourceId,omitempty"`

	// The id of the action to perform.
	ActionID string `json:"actionId,omitempty"`

	// Resource action request inputs
	Inputs map[string]string `json:"inputs,omitempty"`

	// Reason for requesting a day2 operation
	Reason string `json:"reason,omitempty"`
}

// ResourceActionsObservation are the observable fields of a ResourceActions.
type ResourceActionsObservation struct {
	// Name of the resource
	Name *string `json:"name,omitempty"`
	ID   string  `json:"id"`
}

// A ResourceActionsSpec defines the desired state of a ResourceActions.
type ResourceActionsSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       ResourceActionsParameters `json:"forProvider"`
}

// A ResourceActionsStatus represents the observed state of a ResourceActions.
type ResourceActionsStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          ResourceActionsObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// A ResourceActions is an example API type.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,vraprovider}
type ResourceActions struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ResourceActionsSpec   `json:"spec"`
	Status ResourceActionsStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ResourceActionsList contains a list of ResourceActions
type ResourceActionsList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ResourceActions `json:"items"`
}

// ResourceActions type metadata.
var (
	ResourceActionsKind             = reflect.TypeOf(ResourceActions{}).Name()
	ResourceActionsGroupKind        = schema.GroupKind{Group: Group, Kind: ResourceActionsKind}.String()
	ResourceActionsKindAPIVersion   = ResourceActionsKind + "." + SchemeGroupVersion.String()
	ResourceActionsGroupVersionKind = SchemeGroupVersion.WithKind(ResourceActionsKind)
)

func init() {
	SchemeBuilder.Register(&ResourceActions{}, &ResourceActionsList{})
}

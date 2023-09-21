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

// DeploymentActionsParameters are the configurable fields of a DeploymentActions.
type DeploymentActionsParameters struct {
	// Resource deployment id
	DeploymentID string `json:"deploymentId,omitempty"`

	// The id of the action to perform.
	ActionID string `json:"actionId,omitempty"`

	// Resource action request inputs
	Inputs map[string]string `json:"inputs,omitempty"`

	// Reason for requesting a day2 operation
	Reason string `json:"reason,omitempty"`
}

// DeploymentActionsObservation are the observable fields of a DeploymentActions.
type DeploymentActionsObservation struct {
	// Name of the resource
	Name *string `json:"name,omitempty"`
	ID   string  `json:"id"`
}

// A DeploymentActionsSpec defines the desired state of a DeploymentActions.
type DeploymentActionsSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       DeploymentActionsParameters `json:"forProvider"`
}

// A DeploymentActionsStatus represents the observed state of a DeploymentActions.
type DeploymentActionsStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          DeploymentActionsObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// A DeploymentActions is an example API type.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,vraprovider}
type DeploymentActions struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DeploymentActionsSpec   `json:"spec"`
	Status DeploymentActionsStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DeploymentActionsList contains a list of DeploymentActions
type DeploymentActionsList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DeploymentActions `json:"items"`
}

// DeploymentActions type metadata.
var (
	DeploymentActionsKind             = reflect.TypeOf(DeploymentActions{}).Name()
	DeploymentActionsGroupKind        = schema.GroupKind{Group: Group, Kind: DeploymentActionsKind}.String()
	DeploymentActionsKindAPIVersion   = DeploymentActionsKind + "." + SchemeGroupVersion.String()
	DeploymentActionsGroupVersionKind = SchemeGroupVersion.WithKind(DeploymentActionsKind)
)

func init() {
	SchemeBuilder.Register(&DeploymentActions{}, &DeploymentActionsList{})
}

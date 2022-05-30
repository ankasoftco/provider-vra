/*
Copyright 2020 The Crossplane Authors.

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

/*
vRealize Automation Deployment API: https://developer.vmware.com/apis/979
*/

// DeploymentParameters are the configurable fields of a Deployment.
type DeploymentParameters struct {
	// Name of the requested deployment
	// +immutable
	DeploymentName string `json:"deploymentName,omitempty"`

	// Deployment catalog item id
	// +immutable
	CatalogItemID string `json:"catalogItemId,omitempty"`

	// Catalog Item version
	// +optional
	CatalogItemVersion string `json:"catalogItemVersion,omitempty"`

	// Project to be used for the request
	// +immutable
	ProjectID string `json:"projectId,omitempty"`

	// Reason of the deployment
	// +optional
	Reason string `json:"reason,omitempty"`

	// Input parameters for the request. These must be compliant with the schema of the corresponding catalog item
	Inputs map[string]string `json:"inputs,omitempty"`
}

// DeploymentObservation are the observable fields of a Deployment.
type DeploymentObservation struct {
	CreatedAt    string  `json:"createdAt,omitempty"`
	DeploymentID *string `json:"deploymentId,omitempty"`
}

// A DeploymentSpec defines the desired state of a Deployment.
type DeploymentSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       DeploymentParameters `json:"forProvider"`
}

// A DeploymentStatus represents the observed state of a Deployment.
type DeploymentStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          DeploymentObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// A Deployment is an example API type.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,vra}
type Deployment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DeploymentSpec   `json:"spec"`
	Status DeploymentStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DeploymentList contains a list of Deployment
type DeploymentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Deployment `json:"items"`
}

// Deployment type metadata.
var (
	DeploymentKind             = reflect.TypeOf(Deployment{}).Name()
	DeploymentGroupKind        = schema.GroupKind{Group: Group, Kind: DeploymentKind}.String()
	DeploymentKindAPIVersion   = DeploymentKind + "." + SchemeGroupVersion.String()
	DeploymentGroupVersionKind = SchemeGroupVersion.WithKind(DeploymentKind)
)

func init() {
	SchemeBuilder.Register(&Deployment{}, &DeploymentList{})
}

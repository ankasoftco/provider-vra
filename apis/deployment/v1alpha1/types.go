/*
Copyright 2022 The ANKA SOFTWARE Authors.
*/

package v1alpha1

import (
	"reflect"

	"github.com/go-openapi/strfmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

// DeploymentParameters are the configurable fields of a Deployment.
type DeploymentParameters struct {
	// Deployment request count; defaults to 1 if not specified.
	// Maximum: 127
	// Minimum: -128
	// +optional
	BulkRequestCount *int32 `json:"bulkRequestCount,omitempty"`

	// Name of the requested deployment
	// +optional
	DeploymentName string `json:"deploymentName,omitempty"`

	// Input parameters for the request. These must be compliant with the schema of the corresponding catalog item
	// +immutable
	Inputs map[string]string `json:"inputs"` // interface{} `json:"inputs,omitempty"`

	// Project to be used for the request
	// +immutable
	ProjectID string `json:"projectId"`

	// Reason for request
	// +optional
	Reason string `json:"reason,omitempty"`

	// Version of the catalog item. e.g. v2.0
	// +optional
	Version string `json:"version,omitempty"`

	// CatalogItemID Catalog item ID
	// +immutable
	CatalogItemID strfmt.UUID `json:"catalogItemId"`
}

// DeploymentObservation are the observable fields of a Deployment.
type DeploymentObservation struct {
	Name         string      `json:"deploymentName,omitempty"`
	DeploymentID strfmt.UUID `json:"deploymentId,omitempty"`
	Status       string      `json:"status,omitempty"`
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

// A Deployment is a vRA Deployment API type.
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

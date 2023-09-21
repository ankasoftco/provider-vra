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

// BlueprintParameters are the configurable fields of a Blueprint.
type BlueprintParameters struct {

	// Blueprint YAML content
	Content string `json:"content,omitempty"`

	// Content source id
	// Read Only: true
	ContentSourceID string `json:"contentSourceId,omitempty"`

	// Content source path
	// Read Only: true
	ContentSourcePath string `json:"contentSourcePath,omitempty"`

	// Content source type
	// Read Only: true
	ContentSourceType string `json:"contentSourceType,omitempty"`

	// Created by
	// Read Only: true
	CreatedBy string `json:"createdBy,omitempty"`

	// Blueprint description
	Description string `json:"description,omitempty"`

	// Object ID
	// Read Only: true
	ID string `json:"id,omitempty"`

	// Blueprint name
	Name string `json:"name,omitempty"`

	// Org ID
	// Read Only: true
	OrgID string `json:"orgId,omitempty"`

	// Project ID
	ProjectID string `json:"projectId,omitempty"`

	// Project Name
	// Read Only: true
	ProjectName string `json:"projectName,omitempty"`

	// Flag to indicate blueprint can be requested from any project in org
	RequestScopeOrg bool `json:"requestScopeOrg,omitempty"`

	// Blueprint self link
	// Read Only: true
	SelfLink string `json:"selfLink,omitempty"`

	// Blueprint status
	// Read Only: true
	// Enum: [DRAFT VERSIONED RELEASED]
	Status string `json:"status,omitempty"`

	// Total released versions
	// Read Only: true
	TotalReleasedVersions int32 `json:"totalReleasedVersions,omitempty"`

	// Total versions
	// Read Only: true
	TotalVersions int32 `json:"totalVersions,omitempty"`

	// Validation result on update
	// Read Only: true
	Valid *bool `json:"valid,omitempty"`
}

type BlueprintValidationMessage struct {

	// Validation message
	// Read Only: true
	Message string `json:"message,omitempty"`

	// Metadata
	// Read Only: true
	Metadata map[string]string `json:"metadata,omitempty"`

	// Validation path
	// Read Only: true
	Path string `json:"path,omitempty"`

	// Resource name
	// Read Only: true
	ResourceName string `json:"resourceName,omitempty"`

	// Message type
	// Read Only: true
	// Enum: [INFO WARNING ERROR]
	Type string `json:"type,omitempty"`
}

// BlueprintObservation are the observable fields of a Blueprint.
type BlueprintObservation struct {

	// Blueprint YAML content
	Content string `json:"content,omitempty"`

	// Content source type
	// Read Only: true
	ContentSourceType string `json:"contentSourceType,omitempty"`

	// Created by
	// Read Only: true
	CreatedBy string `json:"createdBy,omitempty"`

	// Blueprint description
	Description string `json:"description,omitempty"`

	// Blueprint name
	Name string `json:"name,omitempty"`

	// Project ID
	ProjectID string `json:"projectId,omitempty"`
}

// A BlueprintSpec defines the desired state of a Blueprint.
type BlueprintSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       BlueprintParameters `json:"forProvider"`
}

// A BlueprintStatus represents the observed state of a Blueprint.
type BlueprintStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          BlueprintObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// A Blueprint is an example API type.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,vraprovider}
type Blueprint struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   BlueprintSpec   `json:"spec"`
	Status BlueprintStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// BlueprintList contains a list of Blueprint
type BlueprintList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Blueprint `json:"items"`
}

// Blueprint type metadata.
var (
	BlueprintKind             = reflect.TypeOf(Blueprint{}).Name()
	BlueprintGroupKind        = schema.GroupKind{Group: Group, Kind: BlueprintKind}.String()
	BlueprintKindAPIVersion   = BlueprintKind + "." + SchemeGroupVersion.String()
	BlueprintGroupVersionKind = SchemeGroupVersion.WithKind(BlueprintKind)
)

func init() {
	SchemeBuilder.Register(&Blueprint{}, &BlueprintList{})
}

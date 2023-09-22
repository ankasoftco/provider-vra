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
	"github.com/go-openapi/strfmt"
)

// CatalogSourceParameters are the configurable fields of a CatalogSource.
type CatalogSourceParameters struct {
	// Source custom configuration
	// Required: true
	Config map[string]string `json:"config"`

	// Created By
	CreatedBy string `json:"createdBy,omitempty"`

	// Catalog Source description
	Description string `json:"description,omitempty"`

	// Global flag indicating that all the items can be requested across all projects.
	Global bool `json:"global,omitempty"`

	// Catalog Source name
	// Required: true
	Name *string `json:"name"`

	// Project id where the source belongs
	ProjectID string `json:"projectId,omitempty"`

	// Type of source, e.g. blueprint, CFT... etc
	// Required: true
	TypeID *string `json:"typeId"`
}

// CatalogSourceObservation are the observable fields of a CatalogSource.
type CatalogSourceObservation struct {
	// Source custom configuration
	// Required: true
	Config map[string]string `json:"config"`

	// Created By
	CreatedBy string `json:"createdBy,omitempty"`

	// Catalog Source description
	Description string `json:"description,omitempty"`

	// Global flag indicating that all the items can be requested across all projects.
	Global bool `json:"global,omitempty"`

	// Default Icon Id
	// Format: uuid
	IconID strfmt.UUID `json:"iconId,omitempty"`

	// Catalog Source id
	// Required: true
	// Format: uuid
	ID *strfmt.UUID `json:"id"`

	// Number of items found
	ItemsFound int32 `json:"itemsFound,omitempty"`

	// Number of items imported.
	ItemsImported int32 `json:"itemsImported,omitempty"`

	// Last import error(s)
	LastImportErrors []string `json:"lastImportErrors"`

	// Updated By
	LastUpdatedBy string `json:"lastUpdatedBy,omitempty"`

	// Catalog Source name
	// Required: true
	Name *string `json:"name"`

	// Project id where the source belongs
	ProjectID string `json:"projectId,omitempty"`

	// Type of source, e.g. blueprint, CFT... etc
	// Required: true
	TypeID *string `json:"typeId"`
}

// A CatalogSourceSpec defines the desired state of a CatalogSource.
type CatalogSourceSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       CatalogSourceParameters `json:"forProvider"`
}

// A CatalogSourceStatus represents the observed state of a CatalogSource.
type CatalogSourceStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          CatalogSourceObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// A CatalogSource is an example API type.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,vraprovider}
type CatalogSource struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   CatalogSourceSpec   `json:"spec"`
	Status CatalogSourceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CatalogSourceList contains a list of CatalogSource
type CatalogSourceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CatalogSource `json:"items"`
}

// CatalogSource type metadata.
var (
	CatalogSourceKind             = reflect.TypeOf(CatalogSource{}).Name()
	CatalogSourceGroupKind        = schema.GroupKind{Group: Group, Kind: CatalogSourceKind}.String()
	CatalogSourceKindAPIVersion   = CatalogSourceKind + "." + SchemeGroupVersion.String()
	CatalogSourceGroupVersionKind = SchemeGroupVersion.WithKind(CatalogSourceKind)
)

func init() {
	SchemeBuilder.Register(&CatalogSource{}, &CatalogSourceList{})
}

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

type User struct {

	// The email of the user or name of the group.
	// Example: administrator@vmware.com
	// Required: true
	Email *string `json:"email"`

	// Type of the principal. Currently supported 'user' (default) and 'group'.
	// Example: user
	Type string `json:"type,omitempty"`
}

// ProjectParameters are the configurable fields of a Project.
type ProjectParameters struct {
	// A human-friendly name used as an identifier in APIs that support this option.
	// Required: true - TODO: check required tag
	// +immutable
	Name                  *string `json:"name"`
	MachineNamingTemplate string  `json:"machineNamingTemplate,omitempty"`
	PlacementPolicy       string  `json:"placementPolicy,omitempty"` //DEFAULT, SPREAD, SPREAD BY MEMORY
	OperationTimeout      *int64  `json:"operationTimeout,omitempty"`
	Description           string  `json:"description,omitempty"`
	Administrators        []*User `json:"administrators"`
	SharedResources       bool    `json:"sharedResources,omitempty"`
	Viewers               []*User `json:"viewers"`
	Members               []*User `json:"members"`
}

// ProjectObservation are the observable fields of a Project.
type ProjectObservation struct {
	Name            *string `json:"name"`
	ID              string  `json:"id"`
	Administrators  []*User `json:"administrators"`
	Viewers         []*User `json:"viewers"`
	Members         []*User `json:"members"`
	PlacementPolicy string  `json:"placementPolicy,omitempty"` //DEFAULT, SPREAD, SPREAD BY MEMORY
	//Constraints    map[string][]*struct{} `json:"constraints,omitempty"`
	SharedResources       bool   `json:"sharedResources,omitempty"`
	OperationTimeout      *int64 `json:"operationTimeout,omitempty"`
	MachineNamingTemplate string `json:"machineNamingTemplate,omitempty"`
	Description           string `json:"description,omitempty"`
}

// A ProjectSpec defines the desired state of a Project.
type ProjectSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       ProjectParameters `json:"forProvider"`
}

// A ProjectStatus represents the observed state of a Project.
type ProjectStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          ProjectObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// A Project is an example API type.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,vra}
type Project struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ProjectSpec   `json:"spec"`
	Status ProjectStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ProjectList contains a list of Project
type ProjectList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Project `json:"items"`
}

// Project type metadata.
var (
	ProjectKind             = reflect.TypeOf(Project{}).Name()
	ProjectGroupKind        = schema.GroupKind{Group: Group, Kind: ProjectKind}.String()
	ProjectKindAPIVersion   = ProjectKind + "." + SchemeGroupVersion.String()
	ProjectGroupVersionKind = SchemeGroupVersion.WithKind(ProjectKind)
)

func init() {
	SchemeBuilder.Register(&Project{}, &ProjectList{})
}

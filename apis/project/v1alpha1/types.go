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

// User struct
type User struct {

	// The email of the user or name of the group.
	// Example: administrator@vmware.com
	// Required: true
	Email *string `json:"email"`

	// Type of the principal. Currently supported 'user' (default) and 'group'.
	// Example: user
	Type string `json:"type,omitempty"`
}

// ZoneAssignmentSpecification struct
type ZoneAssignmentSpecification struct {

	// The maximum amount of cpus that can be used by this cloud zone. Default is 0 (unlimited cpu).
	// Example: 2048
	CPULimit int64 `json:"cpuLimit,omitempty"`

	// The maximum number of instances that can be provisioned in this cloud zone. Default is 0 (unlimited instances).
	// Example: 50
	MaxNumberInstances int64 `json:"maxNumberInstances,omitempty"`

	// The maximum amount of memory that can be used by this cloud zone. Default is 0 (unlimited memory).
	// Example: 2048
	MemoryLimitMB int64 `json:"memoryLimitMB,omitempty"`

	// The priority of this zone in the current project. Lower numbers mean higher priority. Default is 0 (highest)
	// Example: 1
	Priority int32 `json:"priority,omitempty"`

	// Defines an upper limit on storage that can be requested from a cloud zone which is part of this project. Default is 0 (unlimited storage). Please note that this feature is supported only for vSphere cloud zones. Not valid for other cloud zone types.
	// Example: 20
	StorageLimitGB int64 `json:"storageLimitGB,omitempty"`

	// The Cloud Zone Id
	// Example: 77ee1
	ZoneID string `json:"zoneId,omitempty"`
}

// Constraint struct
type Constraint struct {

	// An expression of the form "[!]tag-key[:[tag-value]]", used to indicate a constraint match on keys and values of tags.
	//
	// Example: ha:strong
	// Required: true
	Expression *string `json:"expression"`

	// Indicates whether this constraint should be strictly enforced or not.
	// Required: true
	Mandatory *bool `json:"mandatory"`
}

// ProjectParameters are the configurable fields of a Project.
type ProjectParameters struct {
	// A human-friendly name used as an identifier in APIs that support this option.
	// Required: true
	// +immutable
	Name *string `json:"name"`

	// They are not required.
	Administrators               []*User                        `json:"administrators,omitempty"`
	Viewers                      []*User                        `json:"viewers,omitempty"`
	Members                      []*User                        `json:"members,omitempty"`
	MachineNamingTemplate        string                         `json:"machineNamingTemplate,omitempty"`
	PlacementPolicy              string                         `json:"placementPolicy,omitempty"`
	OperationTimeout             *int64                         `json:"operationTimeout,omitempty"`
	Description                  string                         `json:"description,omitempty"`
	SharedResources              bool                           `json:"sharedResources,omitempty"`
	Constraints                  map[string][]Constraint        `json:"constraints,omitempty"`
	CustomProperties             map[string]string              `json:"customProperties,omitempty"`
	ZoneAssignmentConfigurations []*ZoneAssignmentSpecification `json:"zoneAssignmentConfigurations,omitempty"`
}

// ProjectObservation are the observable fields of a Project.
type ProjectObservation struct {
	Name                         *string                        `json:"name,omitempty"`
	ID                           string                         `json:"id"`
	Administrators               []*User                        `json:"administrators,omitempty"`
	Viewers                      []*User                        `json:"viewers,omitempty"`
	Members                      []*User                        `json:"members,omitempty"`
	MachineNamingTemplate        string                         `json:"machineNamingTemplate,omitempty"`
	PlacementPolicy              string                         `json:"placementPolicy,omitempty"`
	OperationTimeout             *int64                         `json:"operationTimeout,omitempty"`
	Description                  string                         `json:"description,omitempty"`
	SharedResources              bool                           `json:"sharedResources,omitempty"`
	Constraints                  map[string][]Constraint        `json:"constraints,omitempty"`
	CustomProperties             map[string]string              `json:"customProperties,omitempty"`
	ZoneAssignmentConfigurations []*ZoneAssignmentSpecification `json:"zoneAssignmentConfigurations,omitempty"`
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

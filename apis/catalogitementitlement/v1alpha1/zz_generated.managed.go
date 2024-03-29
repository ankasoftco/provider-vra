/*
Copyright 2023 Upbound Inc. - ANKASOFT
*/
// Code generated by angryjet. DO NOT EDIT.

package v1alpha1

import xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"

// GetCondition of this CatalogItemEntitlement.
func (mg *CatalogItemEntitlement) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this CatalogItemEntitlement.
func (mg *CatalogItemEntitlement) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetManagementPolicy of this CatalogItemEntitlement.
func (mg *CatalogItemEntitlement) GetManagementPolicy() xpv1.ManagementPolicy {
	return mg.Spec.ManagementPolicy
}

// GetProviderConfigReference of this CatalogItemEntitlement.
func (mg *CatalogItemEntitlement) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this CatalogItemEntitlement.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *CatalogItemEntitlement) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetPublishConnectionDetailsTo of this CatalogItemEntitlement.
func (mg *CatalogItemEntitlement) GetPublishConnectionDetailsTo() *xpv1.PublishConnectionDetailsTo {
	return mg.Spec.PublishConnectionDetailsTo
}

// GetWriteConnectionSecretToReference of this CatalogItemEntitlement.
func (mg *CatalogItemEntitlement) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this CatalogItemEntitlement.
func (mg *CatalogItemEntitlement) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this CatalogItemEntitlement.
func (mg *CatalogItemEntitlement) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetManagementPolicy of this CatalogItemEntitlement.
func (mg *CatalogItemEntitlement) SetManagementPolicy(r xpv1.ManagementPolicy) {
	mg.Spec.ManagementPolicy = r
}

// SetProviderConfigReference of this CatalogItemEntitlement.
func (mg *CatalogItemEntitlement) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this CatalogItemEntitlement.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *CatalogItemEntitlement) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetPublishConnectionDetailsTo of this CatalogItemEntitlement.
func (mg *CatalogItemEntitlement) SetPublishConnectionDetailsTo(r *xpv1.PublishConnectionDetailsTo) {
	mg.Spec.PublishConnectionDetailsTo = r
}

// SetWriteConnectionSecretToReference of this CatalogItemEntitlement.
func (mg *CatalogItemEntitlement) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

/*
Copyright 2022 The ANKA SOFTWARE Authors.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1alpha1

import resource "github.com/crossplane/crossplane-runtime/pkg/resource"

// GetItems of this CatalogItemList.
func (l *CatalogItemList) GetItems() []resource.Managed {
	items := make([]resource.Managed, len(l.Items))
	for i := range l.Items {
		items[i] = &l.Items[i]
	}
	return items
}

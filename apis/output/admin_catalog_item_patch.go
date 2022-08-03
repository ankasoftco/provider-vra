package models 
type AdminCatalogItemPatch struct {

	// Max number of instances that can be requested at a time
	// Maximum: 127
	// Minimum: -128
	BulkRequestLimit *int32 `json:"bulkRequestLimit,omitempty"`

	// form id
	FormID string `json:"formId,omitempty"`

	// icon id
	IconID string `json:"iconId,omitempty"`
}


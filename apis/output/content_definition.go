package models 
type ContentDefinition struct {

	// Description of either the catalog item or the catalog source
	Description string `json:"description,omitempty"`

	// Icon id of associated catalog item (if association is with catalog item)
	// Format: uuid
	IconID strfmt.UUID `json:"iconId,omitempty"`

	// Id of either the catalog source or catalog item.
	// Required: true
	// Format: uuid
	ID *strfmt.UUID `json:"id"`

	// Name of either the catalog item or the catalog source
	Name string `json:"name,omitempty"`

	// Number of items in the associated catalog source
	NumItems int32 `json:"numItems,omitempty"`

	// Catalog source name
	SourceName string `json:"sourceName,omitempty"`

	// Catalog source type
	SourceType string `json:"sourceType,omitempty"`

	// Content definition type
	// Required: true
	Type *string `json:"type"`
}


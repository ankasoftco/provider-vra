package models 
type PropertyDefinition struct {

	// dollar data
	DollarData string `json:"$data,omitempty"`

	// dollar dynamic default
	DollarDynamicDefault string `json:"$dynamicDefault,omitempty"`

	// dollar ref
	DollarRef string `json:"$ref,omitempty"`

	// additional properties
	AdditionalProperties bool `json:"additionalProperties,omitempty"`

	// all of
	AllOf []*PropertyDefinition `json:"allOf"`

	// any of
	AnyOf []*PropertyDefinition `json:"anyOf"`

	// computed
	Computed bool `json:"computed,omitempty"`

	// const
	Const interface{} `json:"const,omitempty"`

	// default
	Default interface{} `json:"default,omitempty"`

	// dependencies
	Dependencies map[string][]string `json:"dependencies,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// encrypted
	Encrypted bool `json:"encrypted,omitempty"`

	// enum
	Enum []interface{} `json:"enum"`

	// format
	Format string `json:"format,omitempty"`

	// ignore case on diff
	IgnoreCaseOnDiff bool `json:"ignoreCaseOnDiff,omitempty"`

	// ignore on update
	IgnoreOnUpdate bool `json:"ignoreOnUpdate,omitempty"`

	// items
	Items *PropertyDefinition `json:"items,omitempty"`

	// max items
	MaxItems int64 `json:"maxItems,omitempty"`

	// max length
	MaxLength int32 `json:"maxLength,omitempty"`

	// max properties
	MaxProperties int32 `json:"maxProperties,omitempty"`

	// maximum
	Maximum int64 `json:"maximum,omitempty"`

	// min items
	MinItems int64 `json:"minItems,omitempty"`

	// min length
	MinLength int32 `json:"minLength,omitempty"`

	// min properties
	MinProperties int32 `json:"minProperties,omitempty"`

	// minimum
	Minimum int64 `json:"minimum,omitempty"`

	// not
	Not *PropertyDefinition `json:"not,omitempty"`

	// one of
	OneOf []*PropertyDefinition `json:"oneOf"`

	// pattern
	Pattern string `json:"pattern,omitempty"`

	// properties
	Properties map[string]PropertyDefinition `json:"properties,omitempty"`

	// read only
	ReadOnly bool `json:"readOnly,omitempty"`

	// recreate on update
	RecreateOnUpdate bool `json:"recreateOnUpdate,omitempty"`

	// required
	Required []string `json:"required"`

	// title
	Title string `json:"title,omitempty"`

	// type
	Type string `json:"type,omitempty"`

	// unique items
	UniqueItems bool `json:"uniqueItems,omitempty"`

	// write only
	WriteOnly bool `json:"writeOnly,omitempty"`
}


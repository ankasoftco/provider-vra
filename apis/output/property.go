package models 
type Property struct {

	// Path that can be used to retrieve single permissible default value
	DollarDynamicDefault string `json:"$dynamicDefault,omitempty"`

	// Path that can be used to retrieve permissible values
	DollarDynamicEnum string `json:"$dynamicEnum,omitempty"`

	// Constant value for the property
	Const interface{} `json:"const,omitempty"`

	// Default value for the property
	Default interface{} `json:"default,omitempty"`

	// Property description
	Description string `json:"description,omitempty"`

	// Flag indicating if property is encrypted
	Encrypted bool `json:"encrypted,omitempty"`

	// Restrict a value to a fixed set of values for the property
	Enum []interface{} `json:"enum"`

	// The expected (and pre-defined) format to be used for string validations. Supported format values include: date-time, email, ip-address, alphanumeric, phone, etc.
	Format string `json:"format,omitempty"`

	// Property items defines a schema for the contents of an array
	Items *Property `json:"items,omitempty"`

	// Maximum number of items for the property
	MaxItems int64 `json:"maxItems,omitempty"`

	// String instance is validated against this keyword for its maximum length
	MaxLength int32 `json:"maxLength,omitempty"`

	// Maximum value for the property
	Maximum int64 `json:"maximum,omitempty"`

	// Minimum number of items for the property
	MinItems int64 `json:"minItems,omitempty"`

	// String instance is validated against this keyword for its minimum length
	MinLength int32 `json:"minLength,omitempty"`

	// Minimum value for the property
	Minimum int64 `json:"minimum,omitempty"`

	// Ensures the given data is valid against only one of the specified schemas
	OneOf []*Property `json:"oneOf"`

	// Regular expression to express constraint for the property
	Pattern string `json:"pattern,omitempty"`

	// Properties within the property
	Properties map[string]Property `json:"properties,omitempty"`

	// Flag indicating if property is read only
	ReadOnly bool `json:"readOnly,omitempty"`

	// Property title
	Title string `json:"title,omitempty"`

	// Property type should be among valid types (number, integer, string, array, boolean, object)
	Type string `json:"type,omitempty"`
}


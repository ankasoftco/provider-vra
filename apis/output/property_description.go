package models 
type PropertyDescription struct {

	// element description
	ElementDescription *PropertyDescription `json:"elementDescription,omitempty"`

	// enum values
	EnumValues []string `json:"enumValues"`

	// example value
	ExampleValue interface{} `json:"exampleValue,omitempty"`

	// field descriptions
	FieldDescriptions map[string]PropertyDescription `json:"fieldDescriptions,omitempty"`

	// indexing options
	IndexingOptions []string `json:"indexingOptions"`

	// kind
	Kind string `json:"kind,omitempty"`

	// property documentation
	PropertyDocumentation string `json:"propertyDocumentation,omitempty"`

	// type name
	// Enum: [LONG STRING BYTES PODO COLLECTION MAP BOOLEAN DOUBLE InternetAddressV4 InternetAddressV6 DATE URI ENUM]
	TypeName string `json:"typeName,omitempty"`

	// usage options
	UsageOptions []string `json:"usageOptions"`
}


package models 
type CustomNamingModel struct {

	// description
	Description string `json:"description,omitempty"`

	// id
	// Format: uuid
	ID strfmt.UUID `json:"id,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// projects
	// Unique: true
	Projects []*CnProjectVo `json:"projects"`

	// templates
	// Unique: true
	Templates []*CnTemplateVo `json:"templates"`
}


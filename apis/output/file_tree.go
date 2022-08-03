package models 
type FileTree struct {

	// The ID of a content source
	// Format: uuid
	RepositoryID strfmt.UUID `json:"repositoryId,omitempty"`

	// A list of directories
	Tree []*FileTreeDirectory `json:"tree"`
}


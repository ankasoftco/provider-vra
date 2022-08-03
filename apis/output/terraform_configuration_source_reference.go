package models 
type TerraformConfigurationSourceReference struct {

	// ID that identifies the commit that corresponds to the desired version of the remote file(s).
	CommitID string `json:"commitId,omitempty"`

	// The ID of the relevant configuration source repository.
	// Format: uuid
	RepositoryID strfmt.UUID `json:"repositoryId,omitempty"`

	// A path to the desired Terraform directory.
	SourceDirectory string `json:"sourceDirectory,omitempty"`
}


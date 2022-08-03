package models 
type CommitDetails struct {

	// author name
	AuthorName string `json:"authorName,omitempty"`

	// comments
	Comments string `json:"comments,omitempty"`

	// commit Id
	CommitID string `json:"commitId,omitempty"`

	// committed date
	// Format: date-time
	CommittedDate strfmt.DateTime `json:"committedDate,omitempty"`

	// committer email
	CommitterEmail string `json:"committerEmail,omitempty"`

	// committer name
	CommitterName string `json:"committerName,omitempty"`
}


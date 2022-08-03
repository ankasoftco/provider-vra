package models 
type Document struct {

	// Document Link
	DownloadURL string `json:"downloadUrl,omitempty"`

	// Document Title
	Title string `json:"title,omitempty"`

	// Document Type
	Type string `json:"type,omitempty"`
}


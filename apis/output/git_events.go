package models 
type gitEvents struct {
	countField int64

	documentsField map[string]GitEvent

	linksField []string

	totalCountField int64
}


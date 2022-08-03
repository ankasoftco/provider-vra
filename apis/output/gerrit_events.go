package models 
type gerritEvents struct {
	countField int64

	documentsField map[string]GerritEvent

	linksField []string

	totalCountField int64
}


package models 
type gerritTriggers struct {
	countField int64

	documentsField map[string]GerritTrigger

	linksField []string

	totalCountField int64
}


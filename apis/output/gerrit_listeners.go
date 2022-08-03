package models 
type gerritListeners struct {
	countField int64

	documentsField map[string]GerritListener

	linksField []string

	totalCountField int64
}


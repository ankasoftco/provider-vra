package models 
type pipelines struct {
	countField int64

	documentsField map[string]Pipeline

	linksField []string

	totalCountField int64
}


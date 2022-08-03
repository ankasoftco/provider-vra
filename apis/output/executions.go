package models 
type executions struct {
	countField int64

	documentsField map[string]Execution

	linksField []string

	totalCountField int64
}


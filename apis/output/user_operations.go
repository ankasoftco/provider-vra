package models 
type userOperations struct {
	countField int64

	documentsField map[string]UserOperation

	linksField []string

	totalCountField int64
}


package models 
type gitWebhooks struct {
	countField int64

	documentsField map[string]GitWebhook

	linksField []string

	totalCountField int64
}


package models 
type customIntegrations struct {
	countField int64

	documentsField map[string]CustomIntegration

	linksField []string

	totalCountField int64
}


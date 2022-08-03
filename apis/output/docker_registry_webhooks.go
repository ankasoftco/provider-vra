package models 
type dockerRegistryWebhooks struct {
	countField int64

	documentsField map[string]DockerRegistryWebHook

	linksField []string

	totalCountField int64
}


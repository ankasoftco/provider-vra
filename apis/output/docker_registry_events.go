package models 
type dockerRegistryEvents struct {
	countField int64

	documentsField map[string]DockerRegistryEvent

	linksField []string

	totalCountField int64
}


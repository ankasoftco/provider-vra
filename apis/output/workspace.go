package models 
type workspace struct {
	autoCloneForTriggerField bool

	cacheField []string

	customPropertiesField map[string]string

	endpointField string

	environmentVariablesField map[string]string

	imageField string

	pathField string

	registryField string

	typeField string
}


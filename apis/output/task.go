package models 
type task struct {
	endpointsField map[string]string

	ignoreFailureField bool

	inputField interface{}

	preConditionField string

	tagsField []string

	typeField string
}


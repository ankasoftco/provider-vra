package models 
type stage struct {
	descriptionField string

	tagsField []string

	taskOrderField []string

	tasksField map[string]Task
}


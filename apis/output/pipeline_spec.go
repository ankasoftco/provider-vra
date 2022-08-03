package models 
type pipelineSpec struct {
	inputMetaField map[string]PropertyMetaData

	concurrencyField int32

	descriptionField *string

	enabledField bool

	iconField string

	inputField interface{}

	nameField *string

	notificationsField NotificationConfiguration

	optionsField []string

	outputField interface{}

	projectField string

	rollbacksField []RollbackConfiguration

	stageOrderField []string

	stagesField map[string]Stage

	starredField PipelineStarredProperty

	stateField string

	tagsField []string

	workspaceField Workspace
}


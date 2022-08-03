package models 
type pipeline struct {
	createTimeInMicrosField int64

	inputMetaField map[string]PropertyMetaData

	linkField string

	projectIdField string

	updateTimeInMicrosField int64

	warningsField []ValidatorResponse

	concurrencyField int32

	createdAtField string

	createdByField string

	descriptionField *string

	enabledField bool

	iconField string

	idField string

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

	updatedAtField string

	updatedByField string

	versionField string

	workspaceField Workspace
}


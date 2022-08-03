package models 
type execution struct {
	createTimeInMicrosField int64

	durationInMicrosField int64

	executedByField string

	inputMetaField map[string]PropertyMetaData

	linkField string

	nestedField bool

	outputMetaField map[string]PropertyMetaData

	pipelineLinkField string

	projectIdField string

	requestTimeInMicrosField int64

	rollbackField bool

	sourceField string

	totalDurationInMicrosField int64

	triggeredByField string

	updateTimeInMicrosField int64

	commentsField string

	createdAtField string

	createdByField string

	descriptionField *string

	iconField string

	idField string

	indexField int64

	inputField interface{}

	nameField *string

	notificationsField []Notification

	outputField interface{}

	projectField string

	reasonField string

	resumedAtField string

	stageOrderField []string

	stagesField map[string]StageExecution

	starredField PipelineStarredProperty

	statusField string

	statusMessageField string

	tagsField []string

	updatedAtField string

	updatedByField string

	versionField string

	workspaceResultsField []*WorkspaceResult
}


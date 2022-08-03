package models 
type stageExecution struct {
	durationInMicrosField int64

	endTimeField int64

	startTimeField int64

	idField string

	nameField *string

	notificationsField []Notification

	rollbackConfigurationField RollbackConfiguration

	rollbackResponseField RollbackResponse

	statusField string

	statusMessageField string

	taskOrderField []string

	tasksField map[string]TaskExecution
}


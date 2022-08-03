package models 
type taskExecution struct {
	durationInMicrosField int64

	endTimeField int64

	startTimeField int64

	endpointsField map[string]string

	failureMessageField string

	idField string

	ignoreFailureField bool

	inputField interface{}

	nameField *string

	notificationsField []Notification

	outputField interface{}

	preConditionField string

	rollbackConfigurationField RollbackConfiguration

	rollbackResponseField RollbackResponse

	statusField string

	statusMessageField string

	typeField string
}


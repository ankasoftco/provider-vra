package models 
type gerritEvent struct {
	createTimeInMicrosField int64

	linkField string

	projectIdField string

	updateTimeInMicrosField int64

	actionTakenField string

	branchField string

	changeNumberField int64

	createdAtField string

	createdByField string

	createdOnField int64

	descriptionField *string

	executionIndexField int64

	executionLinkField string

	executionStatusField string

	gerritProjectField string

	idField string

	listenerField string

	messageField string

	nameField *string

	ownerField string

	patchSetNumberField int64

	pipelineField string

	projectField string

	propertiesField map[string]string

	subjectField string

	triggerField string

	triggerTypeField string

	typeField string

	updatedAtField string

	updatedByField string

	versionField string
}


package models 
type gitEvent struct {
	createTimeInMicrosField int64

	linkField string

	projectIdField string

	updateTimeInMicrosField int64

	commitIdField string

	createdAtField string

	createdByField string

	descriptionField *string

	executionIndexField int64

	executionLinkField string

	executionStatusField string

	httpUrlField string

	idField string

	messageField string

	nameField *string

	ownerField string

	pipelineField string

	projectField string

	repoField string

	serverUrlField *URI

	serverWebhookIdField string

	subjectField string

	targetBranchField string

	timeStampInMicrosField int64

	updatedAtField string

	updatedByField string

	versionField string
}

